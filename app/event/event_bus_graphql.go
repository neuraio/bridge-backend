package event

import (
	"context"
	"errors"
	"fmt"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"reflect"
	"strconv"
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
  blocks(from: %d, to: %d) {
    number
    timestamp
    transactions {
      hash
      logs {
        topics
        data
        account {
          address
        }
      }
    }
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
}

type eventSubscriber interface {
	subscribeEvents(event chan *LogEvent, nextSignal chan struct{})
}

type FetchThroughGraphQL struct {
	networkId                   networkId
	ethClient                   *ethclient.Client
	graphClient                 *graphql.Client
	blockStep                   int           // how many blocks in one graph query
	heightDelay                 int           // how many blocks delay comparing to the latest block height
	fetchInterval               time.Duration // time wait between each single graph query
	synchronizeProgressRecorder synchronizeProgressRecord
}

func NewEventFetchThroughGraphQL(rpcClient *ethclient.Client, graphClient *graphql.Client, blockStep, blockDelay int, networkId networkId, interval time.Duration, dataRecordTransaction dataRecorderTransaction) (*FetchThroughGraphQL, error) {

	height, err := rpcClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	logrus.Infof("the latest block %d of network %d", height, networkId)

	if height < uint64(blockDelay) {
		return nil, errors.New("the latest block height is lower than blocks delay")
	}

	if err := graphClient.Run(context.Background(), graphql.NewRequest(fmt.Sprintf(eventLogGraphQuery, height-uint64(blockDelay), height-uint64(blockDelay-blockStep))), nil); err != nil {
		return nil, err
	}

	eventFetcher := &FetchThroughGraphQL{
		networkId:     networkId,
		ethClient:     rpcClient,
		graphClient:   graphClient,
		blockStep:     blockStep,
		heightDelay:   blockDelay,
		fetchInterval: interval,
	}

	if dataRecordTransaction != nil {
		eventFetcher.synchronizeProgressRecorder = newSynchronizeProgressMysqlRecorder(dataRecordTransaction, networkId)
	}

	return eventFetcher, nil
}

type eventLogGraphQueryAccountModel struct {
	Address string `json:"address"`
}

type eventLogGraphQueryLogModel struct {
	Data    string                         `json:"data"`
	Topics  []string                       `json:"topics"`
	Account eventLogGraphQueryAccountModel `json:"account"`
}

type eventLogGraphQueryTransactionModel struct {
	Hash string                       `json:"hash"`
	Logs []eventLogGraphQueryLogModel `json:"logs"`
}

type eventLogGraphQueryBlockModel struct {
	Number       interface{}                          `json:"number"` // eg: unsigned integer for ethereum however string for bsc
	Timestamp    string                               `json:"timestamp"`
	Transactions []eventLogGraphQueryTransactionModel `json:"transactions"`
}

type eventLogGraphQueryResponseModel struct {
	Blocks []eventLogGraphQueryBlockModel `json:"blocks"`
}

var mutex = new(sync.Mutex)

const logDataMaxLength = 50

// this method will block
func (ef *FetchThroughGraphQL) subscribeEvents(event chan *LogEvent, nextSignal chan struct{}) {
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

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			height, err := ef.ethClient.BlockNumber(ctx)
			if ef.networkId == 80001 {
				chainId, _ := ef.ethClient.ChainID(context.Background())
				nid, _ := ef.ethClient.NetworkID(context.Background())
				fmt.Println("subscribeEvents 80001", height, err, chainId.Uint64(), nid.Uint64())
			}
			cancel()
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
			ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
			err = ef.graphClient.Run(ctx, graphql.NewRequest(fmt.Sprintf(eventLogGraphQuery, fetchHeightFloor, fetchHeightCell)), &graphQLResponse)
			cancel()
			if err != nil {
				logrus.Errorf("subscribeEvents, network: %d, graph query error: %s", ef.networkId, err.Error())
				continue
			}

			for _, block := range graphQLResponse.Blocks {
				var blockNumber uint64
				switch reflect.TypeOf(block.Number).Kind() {
				case reflect.Float64:
					blockNumber = uint64(block.Number.(float64))
				case reflect.String:
					blockNumber, err = strconv.ParseUint(block.Number.(string), 0, 64)
					if err != nil {
						logrus.Errorf("invalid string block height, %s\n", block.Number.(string))
						goto endLoop
					}
				default:
					logrus.Errorf("invalid block height, %v\n", block.Number)
					goto endLoop
				}

				logrus.Infoln("subscribeEvents fetch new block ", blockNumber)

				// avoid updating a row at the same time,otherwise it may be deadlock. between mysql transaction and event handler loop control
				mutex.Lock()

				if err := ef.synchronizeProgressRecorder.startRecord(blockNumber); err != nil {
					logrus.Errorf("subscribeEvents, network: %d, record error: %s", ef.networkId, err.Error())
					mutex.Unlock()
					goto endLoop
				}

				for _, transaction := range block.Transactions {
					for _, eventLog := range transaction.Logs {
						if len(eventLog.Topics) == 0 {
							continue
						}

						if len(eventLog.Data) > logDataMaxLength {
							logrus.Infof("subscribeEvents fetch new event log, hash: %s, address: %s, topics: %v, data: %s...", transaction.Hash, eventLog.Account.Address, eventLog.Topics, eventLog.Data[:logDataMaxLength])
						} else {
							logrus.Infof("subscribeEvents fetch new event log, hash: %s, address: %s, topics: %v, data: %s...", transaction.Hash, eventLog.Account.Address, eventLog.Topics, eventLog.Data)
						}

						logEvent := &LogEvent{
							Topic:           eventLog.Topics[0],
							Data:            eventLog.Data,
							blockNumber:     blockNumber,
							transactionHash: transaction.Hash,
							networkId:       ef.networkId,
						}

						if len(eventLog.Topics) > 1 {
							logEvent.Args = eventLog.Topics[1:]
						}

						event <- logEvent
						// wait for all sql completed
						<-nextSignal
					}
				}

				reason, err := ef.synchronizeProgressRecorder.autoCommitRecord()
				if err != nil {
					logrus.Errorf("data record commit error. block height: %d error message: %s\n", blockNumber, err.Error())
				} else if reason != "" {
					logrus.Errorf("data rollback. block height: %d rollback reason: %s\n", blockNumber, reason)
				}

				mutex.Unlock()
			}

		endLoop:
		}
	}
}

var graphLog = func(s string) {
	logrus.Debugln(s)
}

func registerEventSystem(dataRecordTransactions map[networkId]dataRecorderTransaction, nextSignals map[networkId]chan struct{}, eventLog chan *LogEvent) {

	valutConfig := config.GetVaultConfig()

	// get erc20ContractPairs
	networkIds := make([]networkId, 0)

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

		if chain.Graph == "" {
			eventSubscriber, err = NewEventFetchThroughRpc(rpcClient, chain.BlockStep, chain.BlockDelay, networkId(chain.NetworkId), 3*time.Second, dataRecordTransactionMysql)
			if err != nil {
				logrus.Fatalln(err)
			}
		} else {
			graphClient := graphql.NewClient(chain.Graph)
			if config.Debug() {
				graphClient.Log = graphLog
			}
			eventSubscriber, err = NewEventFetchThroughGraphQL(rpcClient, graphClient, chain.BlockStep, chain.BlockDelay, networkId(chain.NetworkId), 3*time.Second, dataRecordTransactionMysql)
			if err != nil {
				logrus.Fatalln(err)
			}
		}

		signer, err := tools.NewVaultSigner(valutConfig.Endpoint, valutConfig.Role, valutConfig.Secret, chain.NetworkId)
		if err != nil {
			logrus.Fatalln(err)
		}

		if err := registerNodeClients(networkId(chain.NetworkId), chain.BridgeContract721, chain.BridgeContract20, rpcClient, chain.AdminAddress, signer.Signer, signer.SignMintErc20); err != nil {
			logrus.Fatalln(err)
		}

		logrus.Infoln("event subscriber start: ", chain.NetworkId)
		networkIds = append(networkIds, networkId(chain.NetworkId))
		go eventSubscriber.subscribeEvents(eventLog, nextSignal)
	}

	erc20ContractPairsConfiguration := make([]*database.Erc20BridgeContractAddress, 0)
	if err := database.GetMysqlClient().Find(&erc20ContractPairsConfiguration).Error; err != nil {
		panic(err)
	}

	erc20ContractPairMap := make(map[string][]*Erc20ContractAddress, 0)
	for i := range erc20ContractPairsConfiguration {
		if _, found := erc20ContractPairMap[erc20ContractPairsConfiguration[i].Name]; !found {
			erc20ContractPairMap[erc20ContractPairsConfiguration[i].Name] = make([]*Erc20ContractAddress, 0)
		}
		erc20ContractPairMap[erc20ContractPairsConfiguration[i].Name] = append(erc20ContractPairMap[erc20ContractPairsConfiguration[i].Name], &Erc20ContractAddress{
			NetworkId:       networkId(erc20ContractPairsConfiguration[i].NetworkId),
			ContractAddress: strings.ToLower(erc20ContractPairsConfiguration[i].ContractAddress),
		})
	}

	ticker := time.NewTicker(time.Minute)

	go func() {
		for true {
			select {
			case <-ticker.C:
				registerErc20ContractPairs(networkIds, erc20ContractPairMap)
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
