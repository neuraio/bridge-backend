package event

import (
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ApeGame/bridge-backend/app/database"
	bridge "github.com/ApeGame/bridge-backend/app/pkg/node/abi"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	recordsForOnceJob    = 200
	mintEventTopic       = ""
	BridgeEventTopic     = "0xa32c8d97b98adc135b448bc36f8fd4fbdc09c4a7bd832e5e9a0510051a75f89c" //event Sent(address sender, address srcNft, uint256 tokenId, uint256 dstChId, address reciver, address dstNft)
	BridgeBurnErc20Topic = "0x50e22ad3fc6c213f811692f757b38468af04f08c4377a69004aaf21c7f04485b" //event Burned(bytes32 indexed burnId, address sender, address receiver, address token, uint256 amount, uint256 dstChainId, uint256 nonce)

	ZkBridgeErc20Topic = "0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b" //event BridgeEvent (uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
	ZkClaimErc20Topic  = "0x25308c93ceeed162da955b3f7ce3e3f93606579e40fb92029faa9efe27545983" //event ClaimEvent (uint32 index, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)

	zkDepositErc20Topic         = "0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d" // event DepositInitiated (index_topic_1 bytes32 l2DepositTxHash, index_topic_2 address from, index_topic_3 address to, address l1Token, uint256 amount)
	zkWithdrawErc20Topic        = "0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0" // event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
	zkL1MessageSentErc20Topic   = "0x3a36e47291f4201faf137fab081d92295bce2d53be2c6ca68ba82c7faa9ce241" // event L1MessageSent(address indexed _sender, bytes32 indexed _hash, bytes _message)
	zkSyncWithdrawBlockNumTopic = "0x2402307311a4d6604e4e7b4c8a15a7e1213edb39c16a31efa70afb06030d3165"
	zkSyncFinalizeWithdrawTopic = "0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383"
	zkDepositRefundErc20Topic   = "0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60"
)

type Erc20ContractAddress struct {
	NetworkId             networkId
	ContractAddress       string
	MinimumFee            *big.Int
	RollupContractAddress string
	DstNetworkId          networkId
	LDstNetworkId         networkId // l1 l2 network id
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
var abiObject721, abiObject20, abiZkObject20, l1ZkObject20, l1ZkMsgSenderObject20, l2ZkObject20 *abi.ABI

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

	if strings.ToLower(bridgeEvent.OriginAddress.String()) != strings.ToLower(rollupAddress) {
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

	if strings.ToLower(bridgeEvent.OriginAddress.String()) != strings.ToLower(rollupAddress) {
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

var bridgeEventZKDepositErc20Handle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	if event.Topic != zkDepositErc20Topic {
		return errors.New("invalid topic")
	}

	if len(event.Args) < 3 {
		return errors.New("invalid event log lacking of args")
	}

	if len(event.Data) <= 2 {
		return errors.New("invalid event log without data")
	}

	// deserialization
	bridgeEvent := new(bridge.L1ERC20BridgeDepositInitiated)
	hexData, err := hex.DecodeString(event.Data[2:])
	if err != nil {
		return err
	}

	if err := l1ZkObject20.UnpackIntoInterface(bridgeEvent, "DepositInitiated", hexData); err != nil {
		return err
	}

	rollupTokenAddress, dstNetwork := getRollUpTokenAddress(event.networkId)
	if rollupTokenAddress == "" || dstNetwork == 0 {
		logrus.Warnf("[Skip] bridgeEventZKSyncDepositErc20Handle Erc20 Bridge Event Fetched Without Any Contract Pair. Transaction Hash: %s", event.transactionHash)
		return nil
	}
	if strings.ToLower(bridgeEvent.L1Token.String()) != strings.ToLower(rollupTokenAddress) {
		return nil
	}
	recorder := &database.BridgeHistory{
		ProtocolType: database.Erc20,

		SourceNetworkId:       int(event.networkId),
		SourceBlockHeight:     event.blockNumber,
		SourceTransactionHash: event.transactionHash,
		SourceAddress:         common.HexToAddress(event.Args[1]).String(),
		SourceContractAddress: bridgeEvent.L1Token.String(),

		DestinationNetworkId:       int(dstNetwork),
		DestinationTransactionHash: event.Args[0],
		DestinationAddress:         common.HexToAddress(event.Args[2]).String(),
		Erc20Amount:                bridgeEvent.Amount.String(),

		Status: database.NftBridgeDepositing,
	}

	logrus.Debugf("bridgeEventZKSyncDepositErc20Handle %v", recorder)

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}
	if err := mysqlClient.Create(recorder).Error; err != nil {
		logrus.Errorf("bridgeEventZKSyncDepositErc20Handle mysqlClient.Save(recorder) %s", err)
		return err
	}
	return nil
}

var bridgeEventZKWithdrawErc20Handle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	if event.Topic != zkWithdrawErc20Topic {
		return errors.New("invalid topic")
	}

	if len(event.Args) < 3 {
		return errors.New("invalid event log lacking of args")
	}

	if len(event.Data) <= 2 {
		return errors.New("invalid event log without data")
	}

	// deserialization
	bridgeEvent := new(bridge.L2ERC20BridgeWithdrawalInitiated)
	hexData, err := hex.DecodeString(event.Data[2:])
	if err != nil {
		return err
	}

	if err := l2ZkObject20.UnpackIntoInterface(bridgeEvent, "WithdrawalInitiated", hexData); err != nil {
		return err
	}

	rollupTokenAddress, dstNetwork := getRollUpTokenAddress(event.networkId)
	if rollupTokenAddress == "" || dstNetwork == 0 {
		logrus.Warnf("[Skip] bridgeEventZKWithdrawErc20Handle Erc20 Bridge Event Fetched Without Any Contract Pair. Transaction Hash: %s", event.transactionHash)
		return nil
	}
	if strings.ToLower(common.HexToAddress(event.Args[2]).String()) != strings.ToLower(rollupTokenAddress) {
		return nil
	}

	recorder := &database.BridgeHistory{
		ProtocolType: database.Erc20,

		SourceNetworkId:       int(event.networkId),
		SourceBlockHeight:     event.blockNumber,
		SourceTransactionHash: event.transactionHash,
		SourceAddress:         common.HexToAddress(event.Args[0]).String(),
		SourceContractAddress: common.HexToAddress(event.Args[2]).String(),

		DestinationNetworkId: int(dstNetwork),
		DestinationAddress:   common.HexToAddress(event.Args[1]).String(),
		Erc20Amount:          bridgeEvent.Amount.String(),
		Status:               database.NftBridgeWithdrawing,
	}

	logrus.Debugf("bridgeEventZKWithdrawErc20Handle %v", recorder)

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	return mysqlClient.Save(recorder).Error
}

var bridgeEventZKSyncWithdrawBlockNumErc20Handle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	if event.Topic != zkSyncWithdrawBlockNumTopic {
		return errors.New("invalid topic")
	}

	if len(event.Args) < 3 {
		return errors.New("invalid event log lacking of args")
	}

	height := big.NewInt(0).SetBytes(common.FromHex(event.Args[0]))

	recorder := &database.SyncZkProgressRecord{
		CreatedAt:   time.Now(),
		BlockHeight: height.Uint64(),
		NetworkId:   int(event.networkId),
		Type:        database.SyncZkFinalizeWithdrawalHeight,
	}

	logrus.Debugf("bridgeEventZKSyncWithdrawBlockNumErc20Handle %v", recorder)

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	return mysqlClient.Save(recorder).Error
}

var bridgeEventZKSyncFinalizeWithdrawErc20Handle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	if event.Topic != zkSyncFinalizeWithdrawTopic {
		return errors.New("invalid topic")
	}

	if len(event.Args) < 2 {
		return errors.New("invalid event log lacking of args")
	}

	l1Token := common.HexToAddress(event.Args[1]).String()
	rollupTokenAddress, dstNetwork := getRollUpTokenAddress(event.networkId)
	if rollupTokenAddress == "" || dstNetwork == 0 {
		logrus.Warnf("[Skip] bridgeEventZKSyncFinalizeWithdrawErc20Handle Erc20 Bridge Event Fetched Without Any Contract Pair. Transaction Hash: %s", event.transactionHash)
		return nil
	}
	if strings.ToLower(l1Token) != strings.ToLower(rollupTokenAddress) {
		logrus.Errorf("l1Token:%s", l1Token)
		logrus.Errorf("rollupTokenAddress:%s", rollupTokenAddress)
		logrus.Warnf("[Skip] bridgeEventZKSyncFinalizeWithdrawErc20Handle l1Token != rollupTokenAddress. Transaction Hash: %s", event.transactionHash)
		return nil
	}

	client, found := nodeClients[event.networkId]
	if !found {
		logrus.Errorf("bridgeEventZKSyncFinalizeWithdrawErc20Handle nodeClients[event.networkId] error. networkID:%s", event.networkId)
		return fmt.Errorf("client not found: %d", event.networkId)
	}
	tx, _, err := client.rpcClient.TransactionByHash(context.Background(), common.HexToHash(event.transactionHash))
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			logrus.Errorf("bridgeEventZKSyncFinalizeWithdrawErc20Handle TransactionByHash error. transactionHash:%s", event.transactionHash)
			return err
		}
		return err
	}

	// decode input data
	parsed, _ := abi.JSON(strings.NewReader(bridge.L1ERC20BridgeABI))
	// 解码输入数据
	res, err := parsed.Methods["finalizeWithdrawal"].Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		logrus.Errorf("bridgeEventZKSyncFinalizeWithdrawErc20Handle parsed.Methods[\"finalizeWithdrawal\"] error. transactionHash:%s", event.transactionHash)
		return err
	}
	var out FinalizeWithdrawalInput
	parsed.Methods["finalizeWithdrawal"].Inputs.Copy(&out, res)

	//l2BlockNumber == l1_batch_number
	//l2MessageIndex == proof_id
	//l2TxNumberInBlock == l1_batch_tx_index

	extra := &database.BridgeHistoryExtra{}
	if err := database.GetMysqlClient().Model(database.BridgeHistoryExtra{}).
		Where("l1_batch_number = ? and proof_id = ? and l1_batch_tx_index = ?", hexutil.EncodeBig(out.L2BlockNumber),
			out.L2MessageIndex.Uint64(), hexutil.EncodeBig(big.NewInt(int64(out.L2TxNumberInBlock)))).First(&extra).Error; err != nil {
		logrus.Errorf("bridgeEventZKSyncFinalizeWithdrawErc20Handle BridgeHistoryExtra error. transactionHash:%s", event.transactionHash)
		return err
	}

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	return mysqlClient.Model(&database.BridgeHistory{}).Where("id = ?", extra.ID).Updates(map[string]interface{}{"destination_block_height": event.blockNumber, "status": database.NftBridgeSuccess}).Error
}

var bridgeEventZKDepositRefundErc20Handle eventHandlerFunction = func(event *LogEvent, transaction dataRecorderTransaction) error {
	if event.Topic != zkDepositRefundErc20Topic {
		return errors.New("invalid topic")
	}

	if len(event.Args) < 2 {
		return errors.New("invalid event log lacking of args")
	}

	if len(event.Data) <= 2 {
		return errors.New("invalid event log without data")
	}

	// deserialization
	bridgeEvent := new(bridge.L1ERC20BridgeClaimedFailedDeposit)
	hexData, err := hex.DecodeString(event.Data[2:])
	if err != nil {
		return err
	}

	if err := l1ZkObject20.UnpackIntoInterface(bridgeEvent, " ClaimedFailedDeposit", hexData); err != nil {
		return err
	}

	rollupTokenAddress, dstNetwork := getRollUpTokenAddress(event.networkId)
	if rollupTokenAddress == "" || dstNetwork == 0 {
		logrus.Warnf("[Skip] bridgeEventZKDepositRefundErc20Handle Erc20 Bridge Event Fetched Without Any Contract Pair. Transaction Hash: %s", event.transactionHash)
		return nil
	}
	if strings.ToLower(bridgeEvent.L1Token.String()) != strings.ToLower(rollupTokenAddress) {
		logrus.Warnf("[Skip] bridgeEventZKDepositRefundErc20Handle l1Token != rollupTokenAddress. Transaction Hash: %s", event.transactionHash)
		return nil
	}

	client, found := nodeClients[event.networkId]
	if !found {
		logrus.Errorf("bridgeEventZKDepositRefundErc20Handle nodeClients[event.networkId] error. networkID:%d", event.networkId)
		return fmt.Errorf("client not found: %d", event.networkId)
	}
	tx, _, err := client.rpcClient.TransactionByHash(context.Background(), common.HexToHash(event.transactionHash))
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			logrus.Errorf("bridgeEventZKSyncFinalizeWithdrawErc20Handle TransactionByHash error. transactionHash:%s", event.transactionHash)
			return err
		}
		return err
	}

	// decode input data
	parsed, _ := abi.JSON(strings.NewReader(bridge.L1ERC20BridgeABI))
	// 解码输入数据
	res, err := parsed.Methods["finalizeWithdrawal"].Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		logrus.Errorf("bridgeEventZKSyncFinalizeWithdrawErc20Handle parsed.Methods[\"finalizeWithdrawal\"] error. transactionHash:%s", event.transactionHash)
		return err
	}
	var out FinalizeWithdrawalInput
	parsed.Methods["claimFailedDeposit"].Inputs.Copy(&out, res)

	//l2BlockNumber == l1_batch_number
	//l2MessageIndex == proof_id
	//l2TxNumberInBlock == l1_batch_tx_index

	extra := &database.BridgeHistoryExtra{}
	if err := database.GetMysqlClient().Model(database.BridgeHistoryExtra{}).
		Where("l1_batch_number = ? and proof_id = ? and l1_batch_tx_index = ?", hexutil.EncodeBig(out.L2BlockNumber),
			out.L2MessageIndex.Uint64(), hexutil.EncodeBig(big.NewInt(int64(out.L2TxNumberInBlock)))).First(&extra).Error; err != nil {
		logrus.Errorf("bridgeEventZKSyncFinalizeWithdrawErc20Handle BridgeHistoryExtra error. transactionHash:%s", event.transactionHash)
		return err
	}

	mysqlClient, ok := transaction.getRawClient().(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	return mysqlClient.Model(&database.BridgeHistory{}).Where("id = ?", extra.ID).Updates(map[string]interface{}{"destination_block_height": event.blockNumber, "status": database.NftBridgeSuccess}).Error
}

type ClaimFailedDeposit struct {
	L2TxHash *big.Int `json:"_l2TxHash"`
}

type FinalizeWithdrawalInput struct {
	L2BlockNumber     *big.Int `json:"_l2BlockNumber"`
	L2MessageIndex    *big.Int `json:"_l2MessageIndex"`
	L2TxNumberInBlock uint16   `json:"_l2TxNumberInBlock"`
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

func getRollUpTokenAddress(sourceNetwork networkId) (string, networkId) {
	erc20ContractPairsLocker.Lock()
	defer erc20ContractPairsLocker.Unlock()

	if len(erc20ContractPairs) == 0 {
		logrus.Error("getRollUpTokenAddress erc20ContractPairs ==0 ")
		return "", 0
	}
	//logrus.Debugf("getRollUpContractAddress sourceNetwork :%d,  erc20ContractPairs:%+v", sourceNetwork, erc20ContractPairs)
	rollupTokenAddress := ""
	var dstNetworkID networkId = 0
	for _, erc20ContractPair := range erc20ContractPairs {
		for j := range erc20ContractPair {
			if erc20ContractPair[j].NetworkId == sourceNetwork && erc20ContractPair[j].LDstNetworkId > 0 {
				rollupTokenAddress = erc20ContractPair[j].ContractAddress
				dstNetworkID = erc20ContractPair[j].LDstNetworkId
			}
		}
	}
	if rollupTokenAddress == "" || dstNetworkID == 0 {
		return "", 0
	}
	return rollupTokenAddress, dstNetworkID
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
			LDstNetworkId:         networkId(erc20ContractPairsConfiguration[i].LDstNetworkId),
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
				LDstNetworkId:         pair.LDstNetworkId,
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
	l1Erc20B, err := bridge.L1ERC20BridgeMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	l1ZkObject20 = l1Erc20B
	l2Erc20B, err := bridge.L2ERC20BridgeMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	l2ZkObject20 = l2Erc20B

	l1MsgErc20B, err := bridge.L1MessageSenderMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	l1ZkMsgSenderObject20 = l1MsgErc20B

	// register handler
	// registerHandlerFunction(mintEventTopic, mintEventHandle)
	registerHandlerFunction(BridgeEventTopic, bridgeEventHandle)
	registerHandlerFunction(BridgeBurnErc20Topic, bridgeEventBurnErc20Handle)
	registerHandlerFunction(ZkBridgeErc20Topic, bridgeEventZKBridgeErc20Handle)
	registerHandlerFunction(ZkClaimErc20Topic, bridgeEventZKClaimErc20Handle)
	registerHandlerFunction(zkDepositErc20Topic, bridgeEventZKDepositErc20Handle)
	registerHandlerFunction(zkWithdrawErc20Topic, bridgeEventZKWithdrawErc20Handle)
	registerHandlerFunction(zkSyncWithdrawBlockNumTopic, bridgeEventZKSyncWithdrawBlockNumErc20Handle)
	registerHandlerFunction(zkSyncFinalizeWithdrawTopic, bridgeEventZKSyncFinalizeWithdrawErc20Handle)

	registerJobs(cronJobWrapper(time.Second, jobSendNftToken), cronJobWrapper(3*time.Minute, jobSendFtToken),
		cronJobWrapper(5*time.Second, jobSendTransactionResult), cronJobWrapper(5*time.Second, jobRefundToken),
		cronJobWrapper(10*time.Second, jobRefundTransactionResult), cronJobWrapper(10*time.Second, jobZkDepositToken),
		cronJobWrapper(10*time.Second, jobZkWithdrawToken),
	)
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

func jobZkDepositToken(_ context.Context) error {
	bridgeHistories := make([]database.BridgeHistory, 0)
	if err := database.GetMysqlClient().Where("status = ?  and destination_transaction_hash != ''", database.NftBridgeDepositing).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
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

func jobZkWithdrawToken(_ context.Context) error {
	bridgeHistories := make([]database.BridgeHistory, 0)
	if err := database.GetMysqlClient().Where("status = ?", database.NftBridgeWithdrawing).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
		return err
	}
	if len(bridgeHistories) == 0 {
		return nil
	}
	ids := make([]uint, len(bridgeHistories))
	for i, bridgeHistory := range bridgeHistories {
		ids[i] = bridgeHistory.ID
	}
	bhes := make([]*database.BridgeHistoryExtra, 0)

	if err := database.GetMysqlClient().Where("id in (?)", ids).Find(&bhes).Error; err != nil {
		return err
	}
	bhesMap := make(map[uint]*database.BridgeHistoryExtra)
	for _, bhe := range bhes {
		bhesMap[bhe.ID] = bhe
	}
	for _, bridgeHistory := range bridgeHistories {
		extra, ok := bhesMap[bridgeHistory.ID]
		if !ok {
			extra = &database.BridgeHistoryExtra{
				ID: bridgeHistory.ID,
			}
		}
		client, found := nodeClients[networkId(bridgeHistory.SourceNetworkId)]
		if !found {
			return fmt.Errorf("client not found: %d", bridgeHistory.SourceNetworkId)
		}
		if extra.Message == "" {
			receipt, err := client.rpcClient.TransactionReceipt(context.Background(), common.HexToHash(bridgeHistory.SourceTransactionHash))
			if err != nil {
				return err
			}

			for _, log := range receipt.Logs {
				if len(log.Topics) > 0 && log.Topics[0].Hex() == zkL1MessageSentErc20Topic {
					bridgeEvent := new(bridge.L1MessageSenderL1MessageSent)
					if err := l1ZkMsgSenderObject20.UnpackIntoInterface(bridgeEvent, "L1MessageSent", log.Data); err != nil {
						logrus.Errorf("l1ZkMsgSenderObject20.UnpackIntoInterface error. err: %s", err)
						break
					}
					extra.Message = hexutil.Encode(bridgeEvent.Message)
					break
				}
			}
		}
		if extra.Proof == "" || extra.ProofID == "" {
			proof, err := tools.GetProof(client.rpcEndpoint, bridgeHistory.SourceTransactionHash)
			if err != nil {
				logrus.Errorf("tools.GetProof error. err: %s", err)
				continue
			}
			proofMar, err := json.Marshal(proof.Proof)
			if err != nil {
				return err
			}
			extra.Proof = string(proofMar)
			extra.ProofID = fmt.Sprint(proof.ID)
		}
		if extra.L1BatchNumber == "" || extra.L1BatchTxIndex == "" {
			batch, err := tools.GetTxL1Batch(client.rpcEndpoint, bridgeHistory.SourceTransactionHash)
			if err != nil {
				logrus.Errorf("tools.GetTxL1Batch error. err: %s", err)
				continue
			}
			extra.L1BatchNumber = batch.L1BatchNumber
			extra.L1BatchTxIndex = batch.L1BatchTxIndex
		}

		if err := database.GetMysqlClient().Save(&extra).Error; err != nil {
			logrus.Error(err)
		}
		height := &database.SyncZkProgressRecord{}
		if err := database.GetMysqlClient().Model(&database.SyncZkProgressRecord{}).
			Where("network_id = ? and type = ?", bridgeHistory.DestinationNetworkId, database.SyncZkFinalizeWithdrawalHeight).
			First(&height).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return err
			}
		}
		l1BatchNum := big.NewInt(0).SetBytes(common.FromHex(extra.L1BatchNumber)).Uint64()

		logrus.Debugf("extra.Message:%d", l1BatchNum)
		if extra.Message != "" && extra.Proof != "" && extra.ProofID != "" && extra.L1BatchNumber != "" && extra.L1BatchTxIndex != "" && height.BlockHeight >= l1BatchNum {
			bridgeHistory.Status = database.NftBridgeFinalizeWithdrawal
		}

		if err := database.GetMysqlClient().Save(&bridgeHistory).Error; err != nil {
			logrus.Error(err)
		}
	}

	return nil
}

func jobZkDepositRefundToken(_ context.Context) error {
	bridgeHistories := make([]database.BridgeHistory, 0)
	if err := database.GetMysqlClient().Where("status = ?", database.NftBridgeDepositRefund).Limit(recordsForOnceJob).Find(&bridgeHistories).Error; err != nil {
		return err
	}
	ids := make([]uint, len(bridgeHistories))
	for i, bridgeHistory := range bridgeHistories {
		ids[i] = bridgeHistory.ID
	}
	bhes := make([]*database.BridgeHistoryExtra, 0)
	if err := database.GetMysqlClient().Where("id in (?)", ids).Find(&bhes).Error; err != nil {
		return err
	}
	bhesMap := make(map[uint]*database.BridgeHistoryExtra)
	for _, bhe := range bhes {
		bhesMap[bhe.ID] = bhe
	}

	for _, bridgeHistory := range bridgeHistories {
		extra, ok := bhesMap[bridgeHistory.ID]
		if !ok {
			extra = &database.BridgeHistoryExtra{
				ID: bridgeHistory.ID,
			}
		}
		client, found := nodeClients[networkId(bridgeHistory.SourceNetworkId)]
		if !found {
			return fmt.Errorf("client not found: %d", bridgeHistory.SourceNetworkId)
		}

		if extra.Proof == "" || extra.ProofID == "" {
			proof, err := tools.GetProof(client.rpcEndpoint, bridgeHistory.SourceTransactionHash)
			if err != nil {
				logrus.Errorf("tools.GetProof error. err: %s", err)
				continue
			}
			proofMar, err := json.Marshal(proof.Proof)
			if err != nil {
				return err
			}
			extra.Proof = string(proofMar)
			extra.ProofID = fmt.Sprint(proof.ID)
		}
		if extra.L1BatchNumber == "" || extra.L1BatchTxIndex == "" {
			batch, err := tools.GetTxL1Batch(client.rpcEndpoint, bridgeHistory.SourceTransactionHash)
			if err != nil {
				logrus.Errorf("tools.GetTxL1Batch error. err: %s", err)
				continue
			}
			extra.L1BatchNumber = batch.L1BatchNumber
			extra.L1BatchTxIndex = batch.L1BatchTxIndex
		}

		if err := database.GetMysqlClient().Save(&extra).Error; err != nil {
			logrus.Error(err)
		}

		//l2BlockNumber == l1_batch_number
		//l2MessageIndex == proof_id
		//l2TxNumberInBlock == l1_batch_tx_index

		//l2MessageIndex proof_id
		//l2TxNumberInBlock l1_batch_tx_index
		//merkleProof Proof

		//if extra.Proof != "" && extra.ProofID != "" && extra.L1BatchNumber != "" && extra.L1BatchTxIndex != "" {
		//	bridgeHistory.Status = database.NftBridgeFinalizeWithdrawal
		//}

		//if err := database.GetMysqlClient().Save(&bridgeHistory).Error; err != nil {
		//	logrus.Error(err)
		//}
	}

	return nil
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
