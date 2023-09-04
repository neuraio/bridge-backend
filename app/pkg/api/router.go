package api

import (
	"errors"
	"fmt"
	"github.com/ApeGame/bridge-backend/app/config"
	"github.com/ApeGame/bridge-backend/app/event"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/ApeGame/bridge-backend/app/database"

	"github.com/ApeGame/bridge-backend/app/pkg/request"
	"github.com/ApeGame/bridge-backend/app/pkg/response"
	"github.com/ApeGame/bridge-backend/app/pkg/service"

	"github.com/gin-gonic/gin"
)

var bridgePoolLimit = new(big.Int).Mul(big.NewInt(5000), big.NewInt(1e18))

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
	v1.PUT("/history-records/:id", updateHistoryRecords)
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
	MinimumFee      string `json:"minimum_fee"`
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
			MinimumFee:      pair.MinFee,
		})
	}

	bridgeContractPairs721 := make([]*database.Erc721BridgeContractAddress, 0)
	if err := database.GetMysqlClient().Find(&bridgeContractPairs721).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	nft := make([]*BridgePair, 0, len(bridgeContractPairs721))

	for _, pair := range bridgeContractPairs721 {
		nft = append(nft, &BridgePair{
			Network:         pair.NetworkId,
			ContractAddress: pair.Address,
			ContractIcon:    pair.Icon,
			ContractName:    pair.Name,
			BurnMin:         "1",
			BurnMax:         "1",
			MinimumFee:      "0",
		})
	}

	c.JSON(http.StatusOK, response.Ok(gin.H{"ft": ft, "nft": nft}))
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
	failCheck := make([]string, 0)

	erc20s := make([]*database.Erc20BridgeContractAddress, 0)
	if err := database.GetMysqlClient().Find(&erc20s).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	for _, erc20 := range erc20s {
		for _, chain := range config.GetChainCfg() {
			if int(erc20.NetworkId) == chain.NetworkId {
				bridgeContracts := make([]string, 0)
				if chain.BridgeContract20 != "" {
					bridgeContracts = append(bridgeContracts, chain.BridgeContract20)
				}
				for _, contract := range bridgeContracts {
					balance, err := tools.Erc20BalanceOf(chain.RpcUrl, contract, erc20.ContractAddress)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
						return
					}
					if balance.Cmp(bridgePoolLimit) == -1 {
						failCheck = append(failCheck, fmt.Sprintf("bridge %s pool amount %s is lower than %s", contract, balance.String(), bridgePoolLimit.String()))
					}
				}
			}
		}
	}

	spr := make([]*database.SynchronizationProgressRecord, 0)
	if err := database.GetMysqlClient().Find(&spr).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
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

	if err := checkBridgeStatus(); err != nil {
		failCheck = append(failCheck, err.Error())
	}

	if err := checkBridgeStatusFix(); err != nil {
		failCheck = append(failCheck, err.Error())
	}

	if len(failCheck) > 0 {
		c.AbortWithStatusJSON(http.StatusExpectationFailed, failCheck)
		return
	}
	c.JSON(http.StatusOK, "OK")
	return
}

func checkBridgeStatus() error {
	for k, v := range event.PoolBalanceInsufficient {
		if v {
			return fmt.Errorf("erc20 %s bridge pool balance insufficient, please have a check and refill it", k)
		}
	}

	var count int64
	if err := database.GetMysqlClient().Model(new(database.BridgeHistory)).Where("status = ? OR fail_reason != ''", database.NftBridgeFail).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("%d failed transactions exist", count)
	}

	if err := database.GetMysqlClient().Model(new(database.BridgeHistory)).Where("status = ?", database.NftBridgeUndo).Count(&count).Error; err != nil {
		return err
	}
	if count > 50 {
		return fmt.Errorf("%d transactions undo", count)
	}
	return nil
}

func checkBridgeStatusFix() error {
	var count int64
	if err := database.GetMysqlClient().Model(new(database.BridgeHistory)).Where("status in  ?", []int{-1, -2, -3, -4, -5, -6}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("%d transactions need to check", count)
	}
	return nil
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

type historyRecordReq struct {
	Tx string `json:"tx"`
}

func updateHistoryRecords(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, response.Err(response.ErrInvalidParameter))
		return
	}
	req := &historyRecordReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(response.ErrInvalidParameter))
		return
	}
	if err := service.UpdateHistoryRecord(id, req.Tx); err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}
