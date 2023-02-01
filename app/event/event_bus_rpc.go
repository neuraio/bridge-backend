package event

import (
	"context"
	"errors"
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
	blockStep                   int           // how many blocks in one graph query
	heightDelay                 int           // how many blocks delay comparing to the latest block height
	fetchInterval               time.Duration // time wait between each single graph query
	synchronizeProgressRecorder synchronizeProgressRecord
}

func NewEventFetchThroughRpc(rpcClient *ethclient.Client, blockStep, blockDelay int, networkId networkId, interval time.Duration, dataRecordTransaction dataRecorderTransaction) (*eventFetchThroughRpc, error) {

	height, err := rpcClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	logrus.Infof("the latest block %d of network %d", height, networkId)

	if height < uint64(blockDelay) {
		return nil, errors.New("the latest block height is lower than blocks delay")
	}

	return &eventFetchThroughRpc{
		networkId:                   networkId,
		ethClient:                   rpcClient,
		blockStep:                   blockStep,
		heightDelay:                 blockDelay,
		fetchInterval:               interval,
		synchronizeProgressRecorder: newSynchronizeProgressMysqlRecorder(dataRecordTransaction, networkId),
	}, nil
}

func (ef *eventFetchThroughRpc) subscribeEvents(event chan *LogEvent, nextSignal chan struct{}) {
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

			for i := fetchHeightFloor; i <= fetchHeightCell; i++ {
				logs, err := fetchLogWithBlock(ef.ethClient, i)
				if err != nil {
					logrus.Errorf("fetch log with block %d,network: %d", i, ef.networkId)
					goto endLoop
				}

				logrus.Infoln("subscribeEvents fetch new block ", i)

				mutex.Lock()

				if err := ef.synchronizeProgressRecorder.startRecord(i); err != nil {
					logrus.Errorf("subscribeEvents, network: %d, record error: %s\n", ef.networkId, err.Error())
					mutex.Unlock()
					goto endLoop
				}

				for _, log := range logs {
					logEvent := &LogEvent{
						blockNumber:     i,
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

				reason, err := ef.synchronizeProgressRecorder.autoCommitRecord()
				if err != nil {
					logrus.Errorf("data record commit error. block height: %d error message: %s\n", i, err.Error())
				} else if reason != "" {
					logrus.Errorf("data rollback. block height: %d rollback reason: %s\n", i, reason)
				}

				mutex.Unlock()
			}

		endLoop:
		}
	}
}

type LogData struct {
	transactionHash string
	Topics          []common.Hash
	Data            string
}

func fetchLogWithBlock(client *ethclient.Client, blockNumber uint64) ([]*types.Log, error) {
	logrus.Infof("fetchLogWithBlock, block numer: %d", blockNumber)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}

	if block.Transactions().Len() == 0 {
		return nil, nil
	}

	logs := make([]*types.Log, 0)

	for _, transaction := range block.Transactions() {
		receipt, err := client.TransactionReceipt(context.Background(), transaction.Hash())
		if err != nil {
			return nil, err
		}
		logs = append(logs, receipt.Logs...)
	}
	return logs, nil
}
