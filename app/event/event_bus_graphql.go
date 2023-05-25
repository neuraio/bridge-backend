package event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"strings"
	"sync"
	"time"

	"github.com/ApeGame/bridge-backend/app/config"
	"github.com/ApeGame/bridge-backend/app/database"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

const (
	waitForNetworkSynchronizationExponent = 4 // binary descend exponent limit
	waitForNetworkSynchronizationInterval = time.Second
)

const eventLogGraphQuery = `{
  logs(filter: {fromBlock: %d, toBlock: %d, addresses: placeHolderAddresses, topics:[placeHolderTopics]}) {
   account {
	address
  },
	transaction {
	  block {
		number
	  }
	  hash
	}
	topics
	data
  }
}`

type networkId int

type LogEvent struct {
	blockNumber     uint64
	transactionHash string
	networkId       networkId
	Topic           string
	Args            []string
	Data            string
	Address         string
}

type eventSubscriber interface {
	subscribeEvents(event chan *LogEvent, nextSignal, startSignal chan struct{})
}

type FetchThroughGraphQL struct {
	networkId                   networkId
	ethClient                   *ethclient.Client
	graphClient                 *graphql.Client
	logFiltersQuery             string        // including addresses and topics
	blockStep                   int           // how many blocks in one graph query
	heightDelay                 int           // how many blocks delay comparing to the latest block height
	fetchInterval               time.Duration // time wait between each single graph query
	synchronizeProgressRecorder synchronizeProgressRecord
}

func NewEventFetchThroughGraphQL(rpcClient *ethclient.Client, graphClient *graphql.Client, addresses []string, blockStep, blockDelay int, networkId networkId, interval time.Duration, dataRecordTransaction dataRecorderTransaction) (*FetchThroughGraphQL, error) {
	height, err := rpcClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	logrus.Infof("the latest block %d of network %d", height, networkId)

	if height < uint64(blockDelay) {
		return nil, errors.New("the latest block height is lower than blocks delay")
	}

	eventFetcher := &FetchThroughGraphQL{
		networkId:       networkId,
		ethClient:       rpcClient,
		graphClient:     graphClient,
		blockStep:       blockStep,
		heightDelay:     blockDelay,
		fetchInterval:   interval,
		logFiltersQuery: getEventLogGraphQueryWithFilters(addresses),
	}

	if err := graphClient.Run(context.Background(), graphql.NewRequest(fmt.Sprintf(eventFetcher.logFiltersQuery, height-uint64(blockDelay), height-uint64(blockDelay-blockStep))), nil); err != nil {
		return nil, err
	}

	if dataRecordTransaction != nil {
		eventFetcher.synchronizeProgressRecorder = newSynchronizeProgressMysqlRecorder(dataRecordTransaction, networkId)
	}

	return eventFetcher, nil
}

type eventLogGraphQueryBlockModel struct {
	Number uint64 `json:"number"`
}

type eventLogGraphQueryAccountModel struct {
	Address string `json:"address"`
}

type eventLogGraphQueryTransactionModel struct {
	Hash  string                        `json:"hash"`
	Block *eventLogGraphQueryBlockModel `json:"block"`
}

type eventLogGraphQueryLogModel struct {
	Account     *eventLogGraphQueryAccountModel     `json:"account"`
	Transaction *eventLogGraphQueryTransactionModel `json:"transaction"`
	Data        string                              `json:"data"`
	Topics      []string                            `json:"topics"`
}

type eventLogGraphQueryResponseModel struct {
	Logs []*eventLogGraphQueryLogModel `json:"logs"`
}

var mutex = new(sync.Mutex)

const logDataMaxLength = 50

var eventLogGraphQueryWithFilters string

func getEventLogGraphQueryWithFilters(addresses []string) string {
	address, err := json.Marshal(addresses)
	if err != nil {
		logrus.Fatalln(err)
	}

	topic, err := json.Marshal([]string{BridgeBurnErc20Topic, BridgeEventTopic, ZkBridgeErc20Topic,
		ZkClaimErc20Topic,
		zkDepositErc20Topic,
		zkWithdrawErc20Topic,
		zkSyncWithdrawBlockNumTopic,
		zkSyncFinalizeWithdrawTopic})
	if err != nil {
		logrus.Fatalln(err)
	}

	eventLogGraphQueryWithFilters = strings.ReplaceAll(eventLogGraphQuery, "placeHolderTopics", string(topic))
	eventLogGraphQueryWithFilters = strings.ReplaceAll(eventLogGraphQueryWithFilters, "placeHolderAddresses", string(address))

	logrus.Debugln("eventLogGraphQueryAssembly", eventLogGraphQueryWithFilters)

	return eventLogGraphQueryWithFilters
}

// this method will block
func (ef *FetchThroughGraphQL) subscribeEvents(event chan *LogEvent, nextSignal chan struct{}, startSignal chan struct{}) {
	<-startSignal

	ticker := time.NewTicker(ef.fetchInterval)

	var waitForNetworkSynchronization = 0

	for {
		select {
		case <-ticker.C:

			latestHeightRecorded, err := ef.synchronizeProgressRecorder.latestHeight()
			if err != nil {
				logrus.Errorf("subscribeEvents, network: %d, get latest synchronized height record error: %s", ef.networkId, err)
				continue
			}
			fetchHeightFloor := latestHeightRecorded + 1
			fetchHeightCell := fetchHeightFloor + uint64(ef.blockStep) - 1

			height, err := ef.ethClient.BlockNumber(context.Background())
			if err != nil {
				logrus.Errorf("subscribeEvents, network: %d, get latest height from block chain error: %s", ef.networkId, err.Error())
				continue
			}

			if height < fetchHeightCell {
				fetchHeightCell = height
			}

			if fetchHeightCell < fetchHeightFloor {
				waitTime := time.Duration(5<<waitForNetworkSynchronization) * waitForNetworkSynchronizationInterval
				logrus.Warningf("subscribeEvents fetch height cell is lower than floor. wait %s. latest block height: %d, synchronized block height: %d", waitTime.String(), height, fetchHeightFloor)
				time.Sleep(waitTime)
				if waitForNetworkSynchronization < waitForNetworkSynchronizationExponent {
					waitForNetworkSynchronization++
				}
				continue
			}
			waitForNetworkSynchronization = 0

			var graphQLResponse eventLogGraphQueryResponseModel
			err = ef.graphClient.Run(context.Background(), graphql.NewRequest(fmt.Sprintf(ef.logFiltersQuery, fetchHeightFloor, fetchHeightCell)), &graphQLResponse)
			if err != nil {
				logrus.Errorf("subscribeEvents, network: %d, graph query error: %s", ef.networkId, err.Error())
				continue
			}

			// avoid updating a row at the same time,otherwise it may be deadlock. between mysql transaction and event handler loop control
			mutex.Lock()

			if err := ef.synchronizeProgressRecorder.startRecord(); err != nil {
				logrus.Errorf("subscribeEvents, network: %d, record error: %s", ef.networkId, err.Error())
				mutex.Unlock()
				goto endLoop
			}

			for _, eventLog := range graphQLResponse.Logs {

				if len(eventLog.Topics) == 0 {
					continue
				}

				if len(eventLog.Data) > logDataMaxLength {
					logrus.Infof("subscribeEvents fetch new event log, hash: %s, address: %s, topics: %v, data: %s...", eventLog.Transaction.Hash, eventLog.Account.Address, eventLog.Topics, eventLog.Data[:logDataMaxLength])
				} else {
					logrus.Infof("subscribeEvents fetch new event log, hash: %s, address: %s, topics: %v, data: %s...", eventLog.Transaction.Hash, eventLog.Account.Address, eventLog.Topics, eventLog.Data)
				}

				logEvent := &LogEvent{
					Topic:           eventLog.Topics[0],
					Data:            eventLog.Data,
					blockNumber:     eventLog.Transaction.Block.Number,
					transactionHash: eventLog.Transaction.Hash,
					networkId:       ef.networkId,
				}
				if len(eventLog.Topics) > 1 {
					logEvent.Args = eventLog.Topics[1:]
				}

				event <- logEvent
				// wait for all sql completed
				<-nextSignal
			}

			reason, err := ef.synchronizeProgressRecorder.autoCommitRecord(fetchHeightCell)
			if err != nil {
				logrus.Errorf("data record commit error. block height: %d-%d error message: %s", fetchHeightFloor, fetchHeightCell, err.Error())
			} else if reason != "" {
				logrus.Errorf("data rollback. block height: %d-%d rollback reason: %s", fetchHeightFloor, fetchHeightCell, reason)
			} else {
				logrus.Infof("subscribeEvents fetch new logs from: %d to: %d", fetchHeightFloor, fetchHeightCell)
			}

			mutex.Unlock()
		}

	endLoop:
	}
}

var graphLog = func(s string) {
	logrus.Debugln(s)
}

func registerEventSystem(dataRecordTransactions map[networkId]dataRecorderTransaction, nextSignals map[networkId]chan struct{}, eventLog chan *LogEvent) {

	valutConfig := config.GetVaultConfig()

	// get erc20ContractPairs
	networkIds := make([]networkId, 0)
	startSignals := make([]chan struct{}, 0)

	for _, chain := range config.GetChainCfg() {

		blockWithSynchronizationCheck(networkId(chain.NetworkId))

		dataRecordTransactionMysql, nextSignal := newDataRecorderMysqlTransaction(database.GetMysqlClient()), make(chan struct{})
		dataRecordTransactions[networkId(chain.NetworkId)] = dataRecordTransactionMysql
		nextSignals[networkId(chain.NetworkId)] = nextSignal

		var eventSubscriber eventSubscriber

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		rpcClient, err := ethclient.DialContext(ctx, chain.RpcUrl)
		cancel()
		if err != nil {
			logrus.Fatalln(err)
		}
		addresses := make([]string, 0)
		if chain.BridgeContract20 != "" {
			addresses = append(addresses, chain.BridgeContract20)
		}
		if chain.ZKBridgeContract20 != "" {
			addresses = append(addresses, chain.ZKBridgeContract20)
		}
		if chain.BridgeContract721 != "" {
			addresses = append(addresses, chain.BridgeContract721)
		}
		if len(chain.ZKL1L2Contract20) > 0 {
			addresses = append(addresses, chain.ZKL1L2Contract20...)
		}
		if chain.Graph == "" {
			eventSubscriber, err = NewEventFetchThroughRpc(rpcClient, addresses, chain.BlockStep, chain.BlockDelay, networkId(chain.NetworkId), 3*time.Second, dataRecordTransactionMysql)
			if err != nil {
				logrus.Fatalln(err)
			}
		} else {
			graphClient := graphql.NewClient(chain.Graph)
			if config.Debug() {
				graphClient.Log = graphLog
			}
			eventSubscriber, err = NewEventFetchThroughGraphQL(rpcClient, graphClient, addresses, chain.BlockStep, chain.BlockDelay, networkId(chain.NetworkId), 3*time.Second, dataRecordTransactionMysql)
			if err != nil {
				logrus.Fatalln(err)
			}
		}

		signer, err := tools.NewVaultSigner(valutConfig.Endpoint, valutConfig.Role, valutConfig.Secret, chain.NetworkId)
		if err != nil {
			logrus.Fatalln(err)
		}

		if err := registerNodeClients(networkId(chain.NetworkId), chain.BridgeContract721, chain.BridgeContract20, chain.RpcUrl, rpcClient, chain.AdminAddress, signer.Signer, signer.SignMintErc20, signer.SignMintErc20S); err != nil {
			logrus.Fatalln(err)
		}

		logrus.Infoln("event subscriber start: ", chain.NetworkId)
		networkIds = append(networkIds, networkId(chain.NetworkId))
		startSignal := make(chan struct{})
		startSignals = append(startSignals, startSignal)
		go eventSubscriber.subscribeEvents(eventLog, nextSignal, startSignal)
	}

	// init contracts configuration and start handle events
	registerErc20ContractPairs(networkIds)
	for i := 0; i < len(startSignals); i++ {
		close(startSignals[i])
	}

	ticker := time.NewTicker(time.Minute)

	go func() {
		for true {
			select {
			case <-ticker.C:
				registerErc20ContractPairs(networkIds)
			}
		}
	}()
}

func StartEventSystem() {

	logrus.Infoln("StartEventSystem")

	eventLog := make(chan *LogEvent)
	dataRecordTransactions := make(map[networkId]dataRecorderTransaction, 0)
	nextSignals := make(map[networkId]chan struct{})

	// support multiple networks
	registerEventSystem(dataRecordTransactions, nextSignals, eventLog)

	go func() {
		for {
			select {
			case eventLog := <-eventLog:
				transaction, found := dataRecordTransactions[eventLog.networkId]
				if !found {
					logrus.Fatalf("Weird! data recorder transaction client not found. %d", eventLog.networkId)
				}
				signal, found := nextSignals[eventLog.networkId]
				if !found {
					logrus.Fatalf("Weird! data recorder next block signal not found. %d", eventLog.networkId)
				}
				if err := HandleEvent(eventLog.Topic, eventLog, transaction); err != nil {
					transaction.setError(err)
					logrus.Error(err)
				}
				signal <- struct{}{}
			}
		}
	}()
}

func blockWithSynchronizationCheck(id networkId) {
	block := make(chan struct{})

	go func() {

		for {
			time.Sleep(time.Second)
			var count int64
			if err := database.GetMysqlClient().Model(&database.SynchronizationProgressRecord{}).Where("network_id = ? and comment = ?", id, database.SynchronizedLatestHeight).Count(&count).Error; err != nil {
				panic(err)
			}

			if count > 1 {
				panic("blockWithSynchronizationCheck")
			}

			if count == 0 {
				var progress = &database.SynchronizationProgressRecord{
					BlockHeight: 0,
					NetworkId:   int(id),
					Comment:     database.SynchronizedLatestHeight,
				}

				if err := database.GetMysqlClient().Save(progress).Error; err != nil {
					logrus.Error(err)
				}
			}

			if count == 1 {
				progress := new(database.SynchronizationProgressRecord)
				if err := database.GetMysqlClient().Where("network_id = ? and comment = ?", id, database.SynchronizedLatestHeight).Find(progress).Error; err != nil {
					logrus.Error(err)
					continue
				}
				if progress.BlockHeight > 0 {
					break
				}
			}

			logrus.Infof("please config block synchronization start block number %d", id)
		}
		block <- struct{}{}
	}()

	<-block
}
