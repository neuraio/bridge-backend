package event

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
	"time"
)

type eventFetchThroughRpc struct {
	networkId                   networkId
	ethClient                   *ethclient.Client
	logFilter                   ethereum.FilterQuery
	blockStep                   int           // how many blocks in one graph query
	heightDelay                 int           // how many blocks delay comparing to the latest block height
	fetchInterval               time.Duration // time wait between each single graph query
	synchronizeProgressRecorder synchronizeProgressRecord
}

func NewEventFetchThroughRpc(rpcClient *ethclient.Client, addresses []string, blockStep, blockDelay int, networkId networkId, interval time.Duration, dataRecordTransaction dataRecorderTransaction) (*eventFetchThroughRpc, error) {

	height, err := rpcClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	logrus.Infof("the latest block %d of network %d", height, networkId)

	if height < uint64(blockDelay) {
		return nil, errors.New("the latest block height is lower than blocks delay")
	}
	//address := []common.Address{common.HexToAddress(contract20), common.HexToAddress(contract721)}
	logAddresses := make([]common.Address, 0)
	for _, address := range addresses {
		logAddresses = append(logAddresses, common.HexToAddress(address))
	}

	return &eventFetchThroughRpc{
		networkId: networkId,
		ethClient: rpcClient,
		logFilter: ethereum.FilterQuery{
			FromBlock: new(big.Int),
			ToBlock:   new(big.Int),
			Addresses: logAddresses,
			Topics: [][]common.Hash{{common.HexToHash(BridgeEventTopic), common.HexToHash(BridgeBurnErc20Topic),
				common.HexToHash(ZkBridgeErc20Topic), common.HexToHash(ZkClaimErc20Topic),
				common.HexToHash(zkDepositErc20Topic),
				common.HexToHash(zkWithdrawErc20Topic),
				common.HexToHash(zkSyncWithdrawBlockNumTopic),
				common.HexToHash(zkSyncFinalizeWithdrawTopic),
				common.HexToHash(lineaBridgingInitiatedErc20Topic),
				common.HexToHash(lineaMessageClaimErc20Topic),
				common.HexToHash(L1L2MessageHashesAddedToInbox),
				common.HexToHash(blockFinalized),
			},
			},
		},
		blockStep:                   blockStep,
		heightDelay:                 blockDelay,
		fetchInterval:               interval,
		synchronizeProgressRecorder: newSynchronizeProgressMysqlRecorder(dataRecordTransaction, networkId),
	}, nil
}

func (ef *eventFetchThroughRpc) subscribeEvents(event chan *LogEvent, nextSignal chan struct{}, startSignal chan struct{}) {

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

			logrus.Debugf("eventFetchThroughRpc subscribeEvents get latest synchronized height record: %d", latestHeightRecorded)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			height, err := ef.ethClient.BlockNumber(ctx)
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

			mutex.Lock()

			if err := ef.synchronizeProgressRecorder.startRecord(); err != nil {
				logrus.Errorf("subscribeEvents, network: %d, record error: %s\n", ef.networkId, err.Error())
				mutex.Unlock()
				goto endLoop
			}

			ef.logFilter.FromBlock.SetUint64(fetchHeightFloor)
			ef.logFilter.ToBlock.SetUint64(fetchHeightCell)
			logs, err := fetchLogs(ef.ethClient, ef.logFilter)
			if err != nil {
				logrus.Errorf("fetch log with block %d,network: %d", fetchHeightFloor, ef.networkId)
				mutex.Unlock()
				goto endLoop
			}
			logrus.Debugf("ef.logFilter: begin:%d, end:%d, topics:%+v, address:%+v, logLen:%d", ef.logFilter.FromBlock.Uint64(), ef.logFilter.ToBlock.Uint64(), ef.logFilter.Topics, ef.logFilter.Addresses, len(logs))
			for _, log := range logs {
				logEvent := &LogEvent{
					blockNumber:     log.BlockNumber,
					transactionHash: log.TxHash.Hex(),
					networkId:       ef.networkId,
					Data:            hexutil.Encode(log.Data),
				}

				for i := range log.Topics {
					if i == 0 {
						logEvent.Topic = log.Topics[i].Hex()
					} else {
						logEvent.Args = append(logEvent.Args, log.Topics[i].Hex())
					}
				}

				event <- logEvent
				<-nextSignal
			}

			reason, err := ef.synchronizeProgressRecorder.autoCommitRecord(fetchHeightCell)
			if err != nil {
				logrus.Errorf("data record commit error. block height %d-%d error message: %s", fetchHeightFloor, fetchHeightCell, err.Error())
			} else if reason != "" {
				logrus.Errorf("data rollback. block height: %d-%d rollback reason: %s", fetchHeightFloor, fetchHeightCell, reason)
			} else {
				logrus.Infof("subscribeEvents fetch new logs from: %d to: %d  ", fetchHeightFloor, fetchHeightCell)
			}

			mutex.Unlock()
		}
	endLoop:
	}
}

type LogData struct {
	transactionHash string
	Topics          []common.Hash
	Data            string
}

func fetchLogs(client *ethclient.Client, query ethereum.FilterQuery) ([]types.Log, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.FilterLogs(ctx, query)

}
