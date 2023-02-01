package event

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/ApeGame/bridge-backend/app/database"
	bridge "github.com/ApeGame/bridge-backend/app/pkg/node/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

const DefaultBridgeFee = 0.0015

const (
	recordsForOnceJob    = 20
	mintEventTopic       = ""
	bridgeEventTopic     = "0xa32c8d97b98adc135b448bc36f8fd4fbdc09c4a7bd832e5e9a0510051a75f89c" //event Sent(address sender, address srcNft, uint256 tokenId, uint256 dstChId, address reciver, address dstNft)
	bridgeBurnErc20Topic = "0x50e22ad3fc6c213f811692f757b38468af04f08c4377a69004aaf21c7f04485b" //event Burned(bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 dstChainId, uint256 nonce)
)

type Erc20ContractAddress struct {
	NetworkId       networkId
	ContractAddress string
}

var erc20ContractPairs [][]*Erc20ContractAddress
var erc20ContractPairsLocker = new(sync.Mutex)

type eventHandler interface {
	handle(event *LogEvent, transaction dataRecorderTransaction) error
}

type eventHandlerFunction func(event *LogEvent, transaction dataRecorderTransaction) error
type cronJobFunction func(ctx context.Context)

type multiplexer struct {
	handlers map[string]eventHandler
}

var defaultMultiplexer = multiplexer{handlers: make(map[string]eventHandler)}
var abiObject721, abiObject20 *abi.ABI

var _ eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	// double check
	if event.Topic != mintEventTopic {
		return errors.New("invalid topic")
	}
	return nil
}
var bridgeEventHandle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	// double check
	if event.Topic != bridgeEventTopic {
		return errors.New("invalid topic")
	}

	if len(event.Data) <= 2 {
		return errors.New("invalid event log without data")
	}

	// deserialization
	bridgeEvent := new(bridge.BridgeSent)
	hexData, err := hex.DecodeString(event.Data[2:])
	if err != nil {
		return err
	}

	if err := abiObject721.UnpackIntoInterface(bridgeEvent, "Sent", hexData); err != nil {
		return err
	}

	recorder := &database.BridgeHistory{
		ProtocolType: database.Erc721,

		SourceNetworkId:       int(event.networkId),
		SourceContractAddress: bridgeEvent.SrcNft.Hex(),
		SourceBlockHeight:     event.blockNumber,
		SourceTransactionHash: event.transactionHash,
		SourceAddress:         bridgeEvent.Sender.Hex(),

		DestinationNetworkId:       int(bridgeEvent.DstChId.Uint64()),
		DestinationContractAddress: bridgeEvent.DstNft.Hex(),
		DestinationBlockHeight:     0,
		DestinationTransactionHash: "",
		DestinationAddress:         bridgeEvent.Reciver.Hex(),

		TokenId: bridgeEvent.TokenId.Uint64(),
		Status:  database.NftBridgeUndo,
	}

	logrus.Debugf("bridgeEventHandle %v", recorder)

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	return mysqlClient.Save(recorder).Error
}
var bridgeEventBurnErc20Handle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {

	// double check
	if event.Topic != bridgeBurnErc20Topic {
		return errors.New("invalid topic")
	}

	if len(event.Args) < 1 {
		return errors.New("invalid event log lacking of args")
	}

	if len(event.Data) <= 2 {
		return errors.New("invalid event log without data")
	}

	burnId := event.Args[0]

	// deserialization
	bridgeEvent := new(bridge.BridgeErc20Burned)
	hexData, err := hex.DecodeString(event.Data[2:])
	if err != nil {
		return err
	}

	if err := abiObject20.UnpackIntoInterface(bridgeEvent, "Burned", hexData); err != nil {
		return err
	}

	destinationContractAddress := getDestinationContractAddress(event.networkId, networkId(bridgeEvent.DstChainId.Uint64()), bridgeEvent.Token.Hex())
	if destinationContractAddress == "" {
		logrus.Warnf("[Skip] Erc20 Bridge Event Fetched Without Any Contract Pair. Transaction Hash: %s", event.transactionHash)
		return nil
	}

	recorder := &database.BridgeHistory{
		ProtocolType: database.Erc20,
		Erc20BurnId:  burnId,

		SourceNetworkId:       int(event.networkId),
		SourceContractAddress: bridgeEvent.Token.Hex(),
		SourceBlockHeight:     event.blockNumber,
		SourceTransactionHash: event.transactionHash,
		SourceAddress:         bridgeEvent.Sender.Hex(),

		DestinationNetworkId:       int(bridgeEvent.DstChainId.Uint64()),
		DestinationContractAddress: destinationContractAddress,
		DestinationBlockHeight:     0,
		DestinationTransactionHash: "",
		DestinationAddress:         bridgeEvent.Receiver.Hex(),

		Erc20Amount: bridgeEvent.Amount.String(),

		Status: database.NftBridgeUndo,
	}

	logrus.Debugf("erc20 bridgeEventHandle %v", recorder)

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	return mysqlClient.Save(recorder).Error
}

func getDestinationContractAddress(sourceNetwork, destinationNetwork networkId, sourceContractAddress string) string {
	erc20ContractPairsLocker.Lock()
	defer erc20ContractPairsLocker.Unlock()

	if len(erc20ContractPairs) == 0 {
		return ""
	}

	var pairIndex = -1
	for i, erc20ContractPair := range erc20ContractPairs {
		for j := range erc20ContractPair {
			if erc20ContractPair[j].NetworkId == sourceNetwork && strings.ToLower(erc20ContractPair[j].ContractAddress) == strings.ToLower(sourceContractAddress) {
				pairIndex = i
			}
		}
	}

	if pairIndex < 0 {
		return ""
	}
	erc20ContractPair := erc20ContractPairs[pairIndex]
	var destinationContractAddress string

	for i := range erc20ContractPair {
		if erc20ContractPair[i].NetworkId == destinationNetwork {
			destinationContractAddress = erc20ContractPair[i].ContractAddress
		}
	}

	return destinationContractAddress
}

var cronJobs = make([]cronJobFunction, 0)

func (hf eventHandlerFunction) handle(event *LogEvent, transaction dataRecorderTransaction) error {
	return hf(event, transaction)
}

func registerErc20ContractPairs(chainIds []networkId, erc20ContractMap map[string][]*Erc20ContractAddress) {
	erc20ContractPairsLocker.Lock()
	defer erc20ContractPairsLocker.Unlock()

	erc20ContractPairs = make([][]*Erc20ContractAddress, 0, len(erc20ContractMap))

	for name, pairs := range erc20ContractMap {
		erc20ContractPairs = append(erc20ContractPairs, make([]*Erc20ContractAddress, 0, len(pairs)))
		for _, pair := range pairs {
			if !containNetwork(pair.NetworkId, chainIds) {
				logrus.Warnf("%s contains network which has not been registered, registered networks %v", name, chainIds)
				continue
			}
			erc20ContractPairs[len(erc20ContractPairs)-1] = append(erc20ContractPairs[len(erc20ContractPairs)-1], &Erc20ContractAddress{
				NetworkId:       pair.NetworkId,
				ContractAddress: pair.ContractAddress,
			})
		}
	}

	logrus.Infof("refresh erc20 contract pairs successfully %v", erc20ContractPairs)
}

func containNetwork(element networkId, set []networkId) bool {
	for i := range set {
		if element == set[i] {
			return true
		}
	}
	return false
}

// will replace the previous one if this specified topic is duplicated
func registerHandlerFunction(topic string, function func(event *LogEvent, transaction dataRecorderTransaction) error) {
	defaultMultiplexer.handlers[topic] = eventHandlerFunction(function)
}

func registerJobs(cronJob ...cronJobFunction) {
	cronJobs = append(cronJobs, cronJob...)
}

func init() {
	a, e := bridge.BridgeMetaData.GetAbi()
	if e != nil {
		panic(e)
	}
	abiObject721 = a

	b, e := bridge.BridgeErc20MetaData.GetAbi()
	if e != nil {
		panic(e)
	}
	abiObject20 = b

	// register handler
	// registerHandlerFunction(mintEventTopic, mintEventHandle)
	registerHandlerFunction(bridgeEventTopic, bridgeEventHandle)
	registerHandlerFunction(bridgeBurnErc20Topic, bridgeEventBurnErc20Handle)

	registerJobs(cronJobWrapper(time.Second, jobSendToken), cronJobWrapper(5*time.Second, jobSendTransactionResult), cronJobWrapper(5*time.Second, jobRefundToken), cronJobWrapper(10*time.Second, jobRefundTransactionResult))
}

func HandleEvent(topic string, event *LogEvent, transaction dataRecorderTransaction) error {
	handler, found := defaultMultiplexer.handlers[topic]
	if !found {
		// logrus.Debugf("%s handler not registered", topic)
		return nil
	}
	return handler.handle(event, transaction)
}

func StartJobHandler(ctx context.Context) {
	for i := range cronJobs {
		go func(cronJob cronJobFunction) {
			cronJob(ctx)
		}(cronJobs[i])
	}
}

func cronJobWrapper(interval time.Duration, job func(ctx context.Context) error) func(ctx context.Context) {
	return func(ctx context.Context) {
		ticker := time.NewTicker(interval)

		for {
			select {
			case <-ctx.Done():
				logrus.Infoln("cronJob done")
				return
			case <-ticker.C:
				if err := job(ctx); err != nil {
					logrus.Error("cronJob", err)
				}
			}
		}
	}
}

func isAlreadyMinted(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(err.Error(), "token already minted") {
		return true
	}
	return false
}

func jobSendToken(_ context.Context) error {
	bridgeHistories := make([]database.BridgeHistory, 0)
	if err := database.GetMysqlClient().Where("status = ?", database.NftBridgeUndo).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
		return err
	}

	for _, bridgeHistory := range bridgeHistories {
		// to keep idempotent
		bridgeHistory.Status = database.NftBridgePending
		if err := database.GetMysqlClient().Save(&bridgeHistory).Error; err != nil {
			logrus.Error(err)
			continue
		}

		var hash string
		var err error

		switch bridgeHistory.ProtocolType {
		case database.Erc20:
			// send mint order to destination network
			hash, err = MintToken(networkId(bridgeHistory.DestinationNetworkId), networkId(bridgeHistory.SourceNetworkId), bridgeHistory.Erc20BurnId, bridgeHistory.DestinationContractAddress, bridgeHistory.SourceAddress, bridgeHistory.DestinationAddress, bridgeHistory.Erc20Amount, DefaultBridgeFee)
			if err != nil || hash == "" {
				logrus.Errorf("jobSendToken fail: record id: %d, error: %v", bridgeHistory.ID, err)
			}
			if err != nil {
				bridgeHistory.FailReason = err.Error()
			}
		case database.Erc721:
			// send mint order to destination network
			hash, err = sentTo(networkId(bridgeHistory.SourceNetworkId), networkId(bridgeHistory.DestinationNetworkId), bridgeHistory.DestinationContractAddress, bridgeHistory.DestinationAddress, bridgeHistory.TokenId)
			if err != nil || hash == "" {
				logrus.Errorf("jobSendToken fail: record id: %d, error: %v", bridgeHistory.ID, err)
			}
			if isAlreadyMinted(err) {
				bridgeHistory.Status = database.NftBridgeFail
				bridgeHistory.FailReason = "reverted"
			} else if err != nil {
				bridgeHistory.FailReason = err.Error()
			}
		default:
			logrus.Warnf("unsupported protocol %s, bridge id: %d", bridgeHistory.ProtocolType, bridgeHistory.ID)
			continue
		}

		logrus.Debugf("bridge protocol: %s, hash %s, record id: %d", bridgeHistory.ProtocolType, hash, bridgeHistory.ID)
		bridgeHistory.DestinationTransactionHash = hash
		if err := database.GetMysqlClient().Save(&bridgeHistory).Error; err != nil {
			logrus.Error(err)
		}
	}

	return nil
}

func jobSendTransactionResult(_ context.Context) error {
	bridgeHistories := make([]database.BridgeHistory, 0)
	if err := database.GetMysqlClient().Where("status = ? and destination_transaction_hash != ''", database.NftBridgePending).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
		return err
	}
	for _, bridgeHistory := range bridgeHistories {
		status, blockHeight, err := getTransactionResult(context.Background(), networkId(bridgeHistory.DestinationNetworkId), bridgeHistory.DestinationTransactionHash)
		if err != nil {
			logrus.Errorf("record id: %d, error message: %s", bridgeHistory.ID, err.Error())
			continue
		} else if status == unknown {
			continue
		}

		bridgeHistory.DestinationBlockHeight = blockHeight

		if status == success {
			bridgeHistory.Status = database.NftBridgeSuccess
		} else {
			bridgeHistory.Status = database.NftBridgeFail
		}

		if err := database.GetMysqlClient().Save(&bridgeHistory).Error; err != nil {
			logrus.Error(err)
		}
	}
	return nil
}

func jobRefundToken(_ context.Context) error {
	bridgeHistories := make([]database.BridgeHistory, 0)
	if err := database.GetMysqlClient().Where("protocol_type = ? and status = ? and refund_audit != ''", database.Erc721, database.NftBridgeFail).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
		return err
	}

	for _, bridgeHistory := range bridgeHistories {
		// to keep idempotent
		bridgeHistory.Status = database.NftBridgeRefundPending
		if err := database.GetMysqlClient().Save(&bridgeHistory).Error; err != nil {
			logrus.Error(err)
			continue
		}

		hash, err := refund(networkId(bridgeHistory.SourceNetworkId), bridgeHistory.SourceContractAddress, bridgeHistory.SourceAddress, bridgeHistory.TokenId)
		if err != nil || hash == "" {
			logrus.Errorf("jobRefundToken fail: record id: %d, error: %v", bridgeHistory.ID, err)
		}
		if isAlreadyMinted(err) {
			bridgeHistory.Status = database.NftBridgeRefundFail
			bridgeHistory.FailReason = "reverted"
		} else if err != nil {
			bridgeHistory.FailReason = err.Error()
		}

		bridgeHistory.RefundHash = hash

		logrus.Debugf("refund hash %s, record id: %d", hash, bridgeHistory.ID)

		if err := database.GetMysqlClient().Save(&bridgeHistory).Error; err != nil {
			logrus.Error(err)
		}
	}

	return nil
}

func jobRefundTransactionResult(_ context.Context) error {
	bridgeHistories := make([]database.BridgeHistory, 0)
	if err := database.GetMysqlClient().Where("status = ? and refund_hash != ''", database.NftBridgeRefundPending).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
		return err
	}
	for _, bridgeHistory := range bridgeHistories {
		status, blockHeight, err := getTransactionResult(context.Background(), networkId(bridgeHistory.SourceNetworkId), bridgeHistory.RefundHash)
		if err != nil {
			logrus.Errorf("record id: %d, error message: %s", bridgeHistory.ID, err.Error())
			continue
		} else if status == unknown {
			continue
		}

		bridgeHistory.DestinationBlockHeight = blockHeight

		if status == success {
			bridgeHistory.Status = database.NftBridgeRefundSuccess
		} else {
			bridgeHistory.Status = database.NftBridgeRefundFail
		}

		if err := database.GetMysqlClient().Save(&bridgeHistory).Error; err != nil {
			logrus.Error(err)
		}
	}
	return nil
}
