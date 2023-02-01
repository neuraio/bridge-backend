package api

import (
	"errors"
	"fmt"
	"github.com/ApeGame/bridge-backend/app/event"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"

	"github.com/ApeGame/bridge-backend/app/database"

	"github.com/ApeGame/bridge-backend/app/pkg/request"
	"github.com/ApeGame/bridge-backend/app/pkg/response"
	"github.com/ApeGame/bridge-backend/app/pkg/service"

	"github.com/gin-gonic/gin"
)

var EthRpc = make(map[int64]*EthRpcClient)

type EthRpcClient struct {
	Endpoint string
	ChainId  int64
	Client   *ethclient.Client
}

func SetupRouter(e *gin.Engine) {
	v1 := e.Group("/v1/nft")
	v1.GET("/history-records", listHistoryRecords)
	v1.GET("/health", health)
	v1.GET("/ownerOf/:contract_address/:chain_id/:token_id", ownerOf)
	v1.GET("/bridge-pair", listBridgePair)
	v1.GET("/bridge-fee", getBridgeFee)
	v1.GET("/nft-contracts", getNftContracts)
}

type contractInformation struct {
	NetworkId uint64 `json:"network_id"`
	Address   string `json:"address"`
	Symbol    string `json:"symbol"`
	Icon      string `json:"icon"`
}

func getNftContracts(c *gin.Context) {
	contracts := make([]*database.Erc721BridgeContractAddress, 0)
	if err := database.GetMysqlClient().Find(&contracts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	res := make([]*contractInformation, 0, len(contracts))
	for _, c := range contracts {
		res = append(res, &contractInformation{
			NetworkId: c.NetworkId,
			Address:   c.Address,
			Symbol:    c.Symbol,
			Icon:      c.Icon,
		})
	}

	c.JSON(http.StatusOK, response.Ok(res))
}

func getBridgeFee(c *gin.Context) {
	c.JSON(http.StatusOK, response.Ok(gin.H{"ft": event.DefaultBridgeFee, "nft": event.DefaultBridgeFee}))
}

type BridgePair struct {
	Network         uint64 `json:"network"`
	ContractAddress string `json:"contract_address"`
	ContractName    string `json:"contract_name"`
	ContractIcon    string `json:"contract_icon"`
	BurnMin         string `json:"burn_min"`
	BurnMax         string `json:"burn_max"`
}

func listBridgePair(c *gin.Context) {
	bridgeContractPairs := make([]*database.Erc20BridgeContractAddress, 0)
	if err := database.GetMysqlClient().Find(&bridgeContractPairs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	ft := make(map[string][]*BridgePair, 0)
	for _, pair := range bridgeContractPairs {
		if _, found := ft[pair.Name]; !found {
			ft[pair.Name] = make([]*BridgePair, 0)
		}
		ft[pair.Name] = append(ft[pair.Name], &BridgePair{
			Network:         pair.NetworkId,
			ContractAddress: pair.ContractAddress,
			ContractIcon:    pair.ContractIcon,
			ContractName:    pair.ContractName,
			BurnMin:         pair.MinBurn,
			BurnMax:         pair.MaxBurn,
		})
	}

	c.JSON(http.StatusOK, response.Ok(gin.H{"ft": ft, "nft": nil}))
}

func listHistoryRecords(c *gin.Context) {
	f := &request.ListHistoryRecordsFilter{}
	if err := c.BindQuery(f); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrInvalidParameter)
		return
	}
	items, total, recordCount, err := service.ListHistoryRecords(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(gin.H{"items": items, "total": total, "recordCount": recordCount}))
}

type synchronizationProgress struct {
	height      uint64
	lastUpdated time.Time
}

var synchronization = make(map[int]*synchronizationProgress)

func health(c *gin.Context) {
	spr := make([]*database.SynchronizationProgressRecord, 0)
	if err := database.GetMysqlClient().Find(&spr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}
	failCheck := make([]string, 0, len(spr))
	for i := range spr {
		r, f := synchronization[spr[i].NetworkId]
		if !f {
			synchronization[spr[i].NetworkId] = &synchronizationProgress{
				height:      spr[i].BlockHeight,
				lastUpdated: time.Now(),
			}
			continue
		}
		// no new blocks synced in 30 seconds. alert
		if spr[i].BlockHeight <= r.height && time.Now().Sub(r.lastUpdated) > 30*time.Second {
			failCheck = append(failCheck, fmt.Sprintf("no new blocks synced in 30 seconds network: %d. latest health check height recorded: %d, time: %s, the latest synchronization recorded in mysql height: %d ", spr[i].NetworkId, r.height, r.lastUpdated.String(), spr[i].BlockHeight))
		} else {
			r.height = spr[i].BlockHeight
			r.lastUpdated = time.Now()
		}
	}

	if len(failCheck) > 0 {
		c.JSON(http.StatusExpectationFailed, failCheck)
		return
	}
	return
}

func ownerOf(c *gin.Context) {
	contractAddress := c.Param("contract_address")
	chainId := c.Param("chain_id")
	tokenId := c.Param("token_id")

	chain, err := strconv.ParseInt(chainId, 0, 64)
	if err != nil {
		logrus.Error("ownerOf", err)
		c.JSON(http.StatusBadRequest, response.Err(fmt.Errorf("invalid chain id %s", tokenId)))
		return
	}

	rpcClient, found := EthRpc[chain]
	if !found {
		c.JSON(http.StatusBadRequest, response.Err(fmt.Errorf("not support chain %s", chainId)))
		return
	}

	id, err := strconv.ParseUint(tokenId, 0, 64)
	if err != nil {
		logrus.Error("ownerOf", err)
		c.JSON(http.StatusBadRequest, response.Err(fmt.Errorf("invalid token id %s", tokenId)))
		return
	}

	owner, err := tools.Erc721OwnerOfCallPost(rpcClient.Endpoint, contractAddress, id)
	if err != nil {
		if errors.Is(err, tools.TokenNotExistError) {
			c.JSON(http.StatusNoContent, response.Err(err))
			return
		}
		logrus.Error("ownerOf ", contractAddress, " ", chainId, " ", tokenId, " ", err.Error())
		c.JSON(http.StatusBadRequest, response.Err(fmt.Errorf("invalid params")))
		return
	}

	c.JSON(http.StatusOK, response.Ok(map[string]interface{}{"owner address": owner}))
	return
}
