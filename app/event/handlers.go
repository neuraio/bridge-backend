package event

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ApeGame/bridge-backend/app/database"
	bridge "github.com/ApeGame/bridge-backend/app/pkg/node/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"sync"
	"time"
)

const (
	DefaultBridgeFee = 0.002
)

const (
	recordsForOnceJob    = 30
	mintEventTopic       = ""
	BridgeEventTopic     = "0xa32c8d97b98adc135b448bc36f8fd4fbdc09c4a7bd832e5e9a0510051a75f89c" //event Sent(address sender, address srcNft, uint256 tokenId, uint256 dstChId, address reciver, address dstNft)
	BridgeBurnErc20Topic = "0x50e22ad3fc6c213f811692f757b38468af04f08c4377a69004aaf21c7f04485b" //event Burned(bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 dstChainId, uint256 nonce)

	ZkBridgeErc20Topic = "0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b" //event BridgeEvent (uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
	ZkClaimErc20Topic  = "0x25308c93ceeed162da955b3f7ce3e3f93606579e40fb92029faa9efe27545983" //event ClaimEvent (uint32 index, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
)

type Erc20ContractAddress struct {
	NetworkId             networkId
	ContractAddress       string
	MinimumFee            *big.Int
	RollupContractAddress string
	DstNetworkId          networkId
}

var senderLocker = make(map[networkId]map[common.Address]*sync.Mutex)

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
var abiObject721, abiObject20, abiZkObject20 *abi.ABI

var _ eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	// double check
	if event.Topic != mintEventTopic {
		return errors.New("invalid topic")
	}
	return nil
}
var bridgeEventHandle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	// double check
	if event.Topic != BridgeEventTopic {
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
	if event.Topic != BridgeBurnErc20Topic {
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
	bridgeEvent := new(bridge.Bridgeerc20Burned)
	hexData, err := hex.DecodeString(event.Data[2:])
	if err != nil {
		return err
	}

	if err := abiObject20.UnpackIntoInterface(bridgeEvent, "Burned", hexData); err != nil {
		return err
	}

	destinationContractAddress, destinationContractMinimumFee := getDestinationContractAddress(event.networkId, networkId(bridgeEvent.DstChainId.Uint64()), bridgeEvent.Token.Hex())
	if destinationContractAddress == "" {
		logrus.Warnf("[Skip] Erc20 Bridge Event Fetched Without Any Contract Pair. Transaction Hash: %s", event.transactionHash)
		return nil
	}

	fee := big.NewInt(0).Mul(bridgeEvent.Amount, big.NewInt(int64(DefaultBridgeFee*1e8)))
	fee = fee.Div(fee, big.NewInt(1e8))

	if fee.Cmp(destinationContractMinimumFee) == -1 {
		fee = destinationContractMinimumFee
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
		Fee:         fee.String(),

		Status: database.NftBridgeUndo,
	}

	logrus.Debugf("erc20 bridgeEventHandle %v", recorder)

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	return mysqlClient.Save(recorder).Error
}

//bridgeEvent.             claimEvent
//
//depositCount     =        index
//originAddress    =        originAddress
//destinationAddress = destinationAddress
//amount                    = amount

var bridgeEventZKClaimErc20Handle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	// double check
	if event.Topic != ZkClaimErc20Topic {
		return errors.New("invalid topic")
	}

	//if len(event.Args) < 1 {
	//	return errors.New("invalid event log lacking of args")
	//}

	if len(event.Data) <= 2 {
		return errors.New("invalid event log without data")
	}

	// deserialization
	bridgeEvent := new(bridge.PolygonZkEVMBridgeClaimEvent)
	hexData, err := hex.DecodeString(event.Data[2:])
	if err != nil {
		return err
	}

	if err := abiZkObject20.UnpackIntoInterface(bridgeEvent, "ClaimEvent", hexData); err != nil {
		return err
	}

	// validate our contract address
	rollupAddress, dstNetwork := getRollUpContractAddress(event.networkId)
	if rollupAddress == "" || dstNetwork == 0 {
		logrus.Warnf("[Skip] bridgeEventZKClaimErc20Handle Erc20 Bridge Event Fetched Without Any Contract Pair. Transaction Hash: %s", event.transactionHash)
		return nil
	}

	if bridgeEvent.OriginAddress.String() != rollupAddress {
		return nil
	}

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}
	oldRecord := &database.BridgeHistory{}
	if err := mysqlClient.Where(&database.BridgeHistory{
		DepositCount:    uint64(bridgeEvent.Index),
		SourceNetworkId: int(dstNetwork),

		SourceContractAddress: bridgeEvent.OriginAddress.String(),
		DestinationAddress:    bridgeEvent.DestinationAddress.String(),
		Erc20Amount:           bridgeEvent.Amount.String(),
		DestinationNetworkId:  int(event.networkId),
	}).First(&oldRecord).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}
	if oldRecord.ID == 0 {
		recorder := &database.BridgeHistory{
			ProtocolType: database.Erc20,

			//SourceNetworkId:       int(bridgeEvent.OriginNetwork),
			SourceContractAddress: bridgeEvent.OriginAddress.String(),
			//SourceBlockHeight:     event.blockNumber,
			//SourceTransactionHash: event.transactionHash,
			//SourceAddress: bridgeEvent.Sender.Hex(),
			DestinationNetworkId: int(event.networkId),
			//DestinationContractAddress: ,
			DestinationBlockHeight:     event.blockNumber,
			DestinationTransactionHash: event.transactionHash,
			DestinationAddress:         bridgeEvent.DestinationAddress.String(),

			Erc20Amount: bridgeEvent.Amount.String(),

			Status: database.NftBridgeZKDepositSlow,
		}

		logrus.Debugf("erc20 bridgeEventHandle %v", recorder)

		return mysqlClient.Save(recorder).Error
	}
	//recorder := &database.BridgeHistory{
	//	Model: gorm.Model{ID: oldRecord.ID},
	//	//ProtocolType: database.Erc20,
	//
	//	//SourceNetworkId:       int(bridgeEvent.OriginNetwork),
	//	//SourceContractAddress: bridgeEvent.OriginAddress.String(),
	//	//SourceBlockHeight:     event.blockNumber,
	//	//SourceTransactionHash: event.transactionHash,
	//	//SourceAddress: bridgeEvent.Sender.Hex(),
	//	DestinationNetworkId: int(event.networkId),
	//	//DestinationContractAddress: ,
	//	DestinationBlockHeight:     event.blockNumber,
	//	DestinationTransactionHash: event.transactionHash,
	//	DestinationAddress:         bridgeEvent.DestinationAddress.String(),
	//
	//	//Erc20Amount: bridgeEvent.Amount.String(),
	//	Status: database.NftBridgeSuccess,
	//}

	return mysqlClient.Model(&database.BridgeHistory{}).Where("id = ?", oldRecord.ID).Updates(map[string]interface{}{"destination_network_id": int(event.networkId),
		"destination_block_height": event.blockNumber, "destination_transaction_hash": event.transactionHash,
		"status": database.NftBridgeSuccess}).Error
}

var bridgeEventZKBridgeErc20Handle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	// double check
	if event.Topic != ZkBridgeErc20Topic {
		return errors.New("invalid topic")
	}

	//if len(event.Args) < 1 {
	//	return errors.New("bridgeEventZKBridgeErc20Handle invalid event log lacking of args")
	//}

	if len(event.Data) <= 2 {
		return errors.New("invalid event log without data")
	}

	// deserialization
	bridgeEvent := new(bridge.PolygonZkEVMBridgeBridgeEvent)
	hexData, err := hex.DecodeString(event.Data[2:])
	if err != nil {
		return err
	}

	if err := abiZkObject20.UnpackIntoInterface(bridgeEvent, "BridgeEvent", hexData); err != nil {
		return err
	}

	rollupAddress, dstNetwork := getRollUpContractAddress(event.networkId)
	if rollupAddress == "" || dstNetwork == 0 {
		logrus.Warnf("[Skip] bridgeEventZKBridgeErc20Handle Erc20 Bridge Event Fetched Without Any Contract Pair. Transaction Hash: %s", event.transactionHash)
		return nil
	}

	if bridgeEvent.OriginAddress.String() != rollupAddress {
		return nil
	}

	recorder := &database.BridgeHistory{
		ProtocolType: database.Erc20,

		SourceNetworkId:       int(event.networkId),
		SourceContractAddress: bridgeEvent.OriginAddress.String(),
		SourceBlockHeight:     event.blockNumber,
		SourceTransactionHash: event.transactionHash,
		//SourceAddress:         "",

		DestinationNetworkId:       int(dstNetwork),
		DestinationContractAddress: "",
		DestinationBlockHeight:     0,
		DestinationTransactionHash: "",
		DestinationAddress:         bridgeEvent.DestinationAddress.String(),

		Erc20Amount:  bridgeEvent.Amount.String(),
		DepositCount: uint64(bridgeEvent.DepositCount),
		//Fee:         fee.String(),

		Status: database.NftBridgeZKing,
	}

	logrus.Debugf("erc20 bridgeEventHandle :+%v", recorder)

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	return mysqlClient.Save(recorder).Error
}

func getDestinationContractAddress(sourceNetwork, destinationNetwork networkId, sourceContractAddress string) (string, *big.Int) {
	erc20ContractPairsLocker.Lock()
	defer erc20ContractPairsLocker.Unlock()

	if len(erc20ContractPairs) == 0 {
		logrus.Error("getDestinationContractAddress erc20ContractPairs empty ")
		return "", nil
	}
	logrus.Debugf("getDestinationContractAddress sourceNetwork :%d, sourceContractAddress:%s, destinationNetwork:%d, erc20ContractPairs:%+v", sourceNetwork, sourceContractAddress, destinationNetwork, erc20ContractPairs)
	var pairIndex = -1
	for i, erc20ContractPair := range erc20ContractPairs {
		for j := range erc20ContractPair {
			if erc20ContractPair[j].NetworkId == sourceNetwork && strings.ToLower(erc20ContractPair[j].ContractAddress) == strings.ToLower(sourceContractAddress) {
				pairIndex = i
			}
		}
	}

	if pairIndex < 0 {
		return "", nil
	}
	erc20ContractPair := erc20ContractPairs[pairIndex]
	var destinationContractAddress string
	var destinationContractMinimumFee *big.Int

	for i := range erc20ContractPair {
		if erc20ContractPair[i].NetworkId == destinationNetwork {
			destinationContractAddress = erc20ContractPair[i].ContractAddress
			destinationContractMinimumFee = erc20ContractPair[i].MinimumFee
			break
		}
	}

	return destinationContractAddress, destinationContractMinimumFee
}

func getRollUpContractAddress(sourceNetwork networkId) (string, networkId) {
	erc20ContractPairsLocker.Lock()
	defer erc20ContractPairsLocker.Unlock()

	if len(erc20ContractPairs) == 0 {
		logrus.Error("getRollUpContractAddress erc20ContractPairs ==0 ")
		return "", 0
	}
	//logrus.Debugf("getRollUpContractAddress sourceNetwork :%d,  erc20ContractPairs:%+v", sourceNetwork, erc20ContractPairs)
	rollupContractAddress := ""
	var dstNetworkID networkId = 0
	for _, erc20ContractPair := range erc20ContractPairs {
		for j := range erc20ContractPair {
			if erc20ContractPair[j].NetworkId == sourceNetwork && erc20ContractPair[j].RollupContractAddress != "" {
				rollupContractAddress = erc20ContractPair[j].RollupContractAddress
				dstNetworkID = erc20ContractPair[j].DstNetworkId
			}
		}
	}
	if rollupContractAddress == "" || dstNetworkID == 0 {
		return "", 0
	}
	return rollupContractAddress, dstNetworkID
}

var cronJobs = make([]cronJobFunction, 0)

func (hf eventHandlerFunction) handle(event *LogEvent, transaction dataRecorderTransaction) error {
	return hf(event, transaction)
}

func registerErc20ContractPairs(chainIds []networkId) {
	erc20ContractPairsLocker.Lock()
	defer erc20ContractPairsLocker.Unlock()

	erc20ContractPairsConfiguration := make([]*database.Erc20BridgeContractAddress, 0)
	if err := database.GetMysqlClient().Find(&erc20ContractPairsConfiguration).Error; err != nil {
		logrus.Error(err)
		return
	}

	erc20ContractPairMap := make(map[string][]*Erc20ContractAddress, 0)
	for i := range erc20ContractPairsConfiguration {
		if _, found := erc20ContractPairMap[erc20ContractPairsConfiguration[i].Name]; !found {
			erc20ContractPairMap[erc20ContractPairsConfiguration[i].Name] = make([]*Erc20ContractAddress, 0)
		}

		minimumFee := big.NewInt(0)
		var ok bool

		minimumFee, ok = new(big.Int).SetString(erc20ContractPairsConfiguration[i].MinFee, 0)
		if !ok {
			logrus.Errorf("invalid minimum fee %s", erc20ContractPairsConfiguration[i].MinFee)
		}
		erc20ContractPairMap[erc20ContractPairsConfiguration[i].Name] = append(erc20ContractPairMap[erc20ContractPairsConfiguration[i].Name], &Erc20ContractAddress{
			NetworkId:             networkId(erc20ContractPairsConfiguration[i].NetworkId),
			ContractAddress:       strings.ToLower(erc20ContractPairsConfiguration[i].ContractAddress),
			MinimumFee:            minimumFee,
			RollupContractAddress: erc20ContractPairsConfiguration[i].RollupContractAddress,
			DstNetworkId:          networkId(erc20ContractPairsConfiguration[i].DstNetworkId),
		})
	}

	erc20ContractPairs = make([][]*Erc20ContractAddress, 0)

	for name, pairs := range erc20ContractPairMap {
		erc20ContractPairs = append(erc20ContractPairs, make([]*Erc20ContractAddress, 0, len(pairs)))
		for _, pair := range pairs {
			if !containNetwork(pair.NetworkId, chainIds) {
				logrus.Warnf("%s contains network which has not been registered, registered networks %v", name, chainIds)
				continue
			}
			erc20ContractPairs[len(erc20ContractPairs)-1] = append(erc20ContractPairs[len(erc20ContractPairs)-1], &Erc20ContractAddress{
				NetworkId:             pair.NetworkId,
				ContractAddress:       pair.ContractAddress,
				MinimumFee:            pair.MinimumFee,
				RollupContractAddress: pair.RollupContractAddress,
				DstNetworkId:          pair.DstNetworkId,
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

	b, e := bridge.Bridgeerc20MetaData.GetAbi()
	if e != nil {
		panic(e)
	}
	abiObject20 = b
	z, e := bridge.PolygonZkEVMBridgeMetaData.GetAbi()
	if e != nil {
		panic(e)
	}
	abiZkObject20 = z
	// register handler
	// registerHandlerFunction(mintEventTopic, mintEventHandle)
	registerHandlerFunction(BridgeEventTopic, bridgeEventHandle)
	registerHandlerFunction(BridgeBurnErc20Topic, bridgeEventBurnErc20Handle)
	registerHandlerFunction(ZkBridgeErc20Topic, bridgeEventZKBridgeErc20Handle)
	registerHandlerFunction(ZkClaimErc20Topic, bridgeEventZKClaimErc20Handle)

	registerJobs(cronJobWrapper(time.Second, jobSendNftToken), cronJobWrapper(10*time.Second, jobSendFtToken), cronJobWrapper(5*time.Second, jobSendTransactionResult), cronJobWrapper(5*time.Second, jobRefundToken), cronJobWrapper(10*time.Second, jobRefundTransactionResult))
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

func jobSendFtToken(_ context.Context) error {

	rows, err := database.GetMysqlClient().Raw("SELECT count(id),destination_network_id,status FROM bridge_histories WHERE (status = ? OR status = ?) AND protocol_type = ? GROUP BY destination_network_id,status ORDER BY status", database.NftBridgePending, database.NftBridgeUndo, database.Erc20).Rows()
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(rows)

	networks := make(map[int]struct{}, 0)

	for rows.Next() {
		var count int
		var network int
		var status int
		if err := rows.Scan(&count, &network, &status); err != nil {
			logrus.Error(err)
			continue
		}
		if count > 0 {
			switch database.BridgeStatus(status) {
			case database.NftBridgeUndo:
				networks[network] = struct{}{}
			case database.NftBridgePending:
				delete(networks, network)
			}
			logrus.Warnf("Network %d Exist Pending Transactions: %d", network, count)
			continue
		}
	}
	logrus.Debugf("jobSendFtToken networks len:%d", len(networks))
	for network := range networks {
		bridgeHistories := make([]*database.BridgeHistory, 0)
		if err := database.GetMysqlClient().Where("status = ? AND protocol_type = ? AND destination_network_id = ?", database.NftBridgeUndo, database.Erc20, network).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
			return err
		}
		if len(bridgeHistories) == 0 {
			continue
		}

		bridges := make([]*ercBridge, 0, len(bridgeHistories))
		marks, ids := make([]string, 0, len(bridgeHistories)), make([]interface{}, 0, len(bridgeHistories))
		for _, bh := range bridgeHistories {
			bridges = append(bridges, &ercBridge{
				sourceNetworkId:            networkId(bh.SourceNetworkId),
				burnId:                     bh.Erc20BurnId,
				destinationContractAddress: bh.DestinationContractAddress,
				senderAddress:              bh.SourceAddress,
				receiverAddress:            bh.DestinationAddress,
				amountS:                    bh.Erc20Amount,
				feeS:                       bh.Fee,
			})
			ids = append(ids, bh.ID)
			marks = append(marks, "?")
		}

		s := fmt.Sprintf("UPDATE bridge_histories SET status = %d WHERE id in (%s)", database.NftBridgePending, strings.Join(marks, ","))
		if err := database.GetMysqlClient().Exec(s, ids...).Error; err != nil {
			logrus.Error(err)
			continue
		}

		hash, err := MintTokens(networkId(network), bridges)
		if err != nil || hash == "" {
			logrus.Errorf("jobSendToken fail: network id: %d,record ids: %v, error: %v", network, ids, err)
		}
		if errors.Is(err, PoolBalanceInsufficientError) {
			s = fmt.Sprintf("UPDATE bridge_histories SET status = %d WHERE id in (%s)", database.NftBridgeUndo, strings.Join(marks, ","))
			if err := database.GetMysqlClient().Exec(s, ids...).Error; err != nil {
				logrus.Error(err)
			}
			continue
		}
		if err != nil {
			logrus.Warnf("bridge fail. error: %s, protocol: erc20, hash %s, record ids: %v", err.Error(), hash, ids)
			s = fmt.Sprintf("UPDATE bridge_histories SET destination_transaction_hash = '%s',fail_reason = '%s' WHERE id in (%s)", hash, err.Error(), strings.Join(marks, ","))
		} else {
			logrus.Infof("bridge success protocol: erc20, hash %s, record ids: %v", hash, ids)
			s = fmt.Sprintf("UPDATE bridge_histories SET destination_transaction_hash = '%s' WHERE id in (%s)", hash, strings.Join(marks, ","))
		}

		if err := database.GetMysqlClient().Exec(s, ids...).Error; err != nil {
			logrus.Error(err)
			continue
		}
	}

	return nil
}

func jobSendNftToken(_ context.Context) error {
	bridgeHistories := make([]database.BridgeHistory, 0)
	if err := database.GetMysqlClient().Where("status = ? AND protocol_type = ?", database.NftBridgeUndo, database.Erc721).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
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
