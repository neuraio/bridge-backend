package event

import (
	"context"
	"errors"
	"fmt"
	bridge "github.com/ApeGame/bridge-backend/app/pkg/node/abi"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"sync"
	"time"
)

var PoolBalanceInsufficient = make(map[string]bool, 0)
var PoolBalanceInsufficientError = errors.New("pool balance insufficient")

const defaultGas = 400000

type erc20MintSignFn func(address, bridge common.Address, sender, receiver, token string, burnId [32]byte, amount, fee, sourceNetworkId, destinationNetworkId *big.Int) (string, error)
type erc20MultipleMintSignFn func(address, bridge common.Address, destinationNetworkId *big.Int, bridges []bridge.BridgeMintReq) ([]string, error)

type contractCaller struct {
	rpcEndpoint                 string
	rpcClient                   *ethclient.Client
	contractCallerErc721        *bridge.BridgeTransactor
	contractCallerErc20         *bridge.Bridgeerc20Transactor
	authenticator               *bind.TransactOpts
	bridgeContractAddressErc721 common.Address
	bridgeContractAddressErc20  common.Address
	erc20MintSigner             erc20MintSignFn
	erc20MultipleMintSigner     erc20MultipleMintSignFn
}

type receiptStatus int

const (
	unknown = iota
	success
	fail
)

var nodeClients = make(map[networkId]*contractCaller)

func registerNodeClients(networkId networkId, bridge721ContractAddress, bridge20ContractAddress, rpcEndpoint string, client *ethclient.Client, address string, signer bind.SignerFn, erc20MintSigner erc20MintSignFn, erc20MultipleMintSigner erc20MultipleMintSignFn) error {

	if _, found := senderLocker[networkId]; !found {
		senderLocker[networkId] = make(map[common.Address]*sync.Mutex)
	}
	senderLocker[networkId][common.HexToAddress(address)] = new(sync.Mutex)

	opt := tools.GenerateTransactorOpts(common.HexToAddress(address), signer)

	opt.GasLimit = defaultGas
	opt.Signer = signer

	caller721, err := bridge.NewBridgeTransactor(common.HexToAddress(bridge721ContractAddress), client)
	if err != nil {
		return err
	}

	caller20, err := bridge.NewBridgeerc20Transactor(common.HexToAddress(bridge20ContractAddress), client)
	if err != nil {
		return err
	}

	nodeClients[networkId] = &contractCaller{
		rpcEndpoint:                 rpcEndpoint,
		rpcClient:                   client,
		authenticator:               opt,
		contractCallerErc721:        caller721,
		contractCallerErc20:         caller20,
		bridgeContractAddressErc721: common.HexToAddress(bridge721ContractAddress),
		bridgeContractAddressErc20:  common.HexToAddress(bridge20ContractAddress),
		erc20MintSigner:             erc20MintSigner,
		erc20MultipleMintSigner:     erc20MultipleMintSigner,
	}
	return nil
}

type ercBridge struct {
	sourceNetworkId            networkId
	burnId                     string
	destinationContractAddress string
	senderAddress              string
	receiverAddress            string
	amountS                    string
	feeS                       string
}

func MintTokens(destinationNetworkId networkId, bridges []*ercBridge) (string, error) {
	client, found := nodeClients[destinationNetworkId]
	if !found {
		return "", fmt.Errorf("client not found: %d", destinationNetworkId)
	}

	reqs := make([]bridge.BridgeMintReq, 0, len(bridges))
	type T struct {
		Amount     *big.Int
		Fee        *big.Int
		PoolAmount *big.Int
	}
	m := make(map[common.Address]*T, 0)

	for _, b := range bridges {
		if _, f := m[common.HexToAddress(b.destinationContractAddress)]; !f {
			balance, err := tools.Erc20BalanceOf(client.rpcEndpoint, client.bridgeContractAddressErc20.Hex(), b.destinationContractAddress)
			if err != nil {
				return "", err
			}
			m[common.HexToAddress(b.destinationContractAddress)] = &T{
				Amount:     big.NewInt(0),
				Fee:        big.NewInt(0),
				PoolAmount: balance,
			}
		}

		amountI, ok := new(big.Int).SetString(b.amountS, 0)
		if !ok {
			return "", fmt.Errorf("bridge burn id: %s amount %s is invalid", b.burnId, b.amountS)
		}
		fee, ok := big.NewInt(0).SetString(b.feeS, 0)
		if !ok {
			return "", fmt.Errorf("invalid fee %s", b.feeS)
		}

		m[common.HexToAddress(b.destinationContractAddress)].Amount.Add(m[common.HexToAddress(b.destinationContractAddress)].Amount, amountI)
		m[common.HexToAddress(b.destinationContractAddress)].Fee.Add(m[common.HexToAddress(b.destinationContractAddress)].Fee, fee)

		if m[common.HexToAddress(b.destinationContractAddress)].Amount.Cmp(m[common.HexToAddress(b.destinationContractAddress)].PoolAmount) == 1 {
			PoolBalanceInsufficient[b.destinationContractAddress] = true
			return "", PoolBalanceInsufficientError
		} else {
			PoolBalanceInsufficient[b.destinationContractAddress] = false
		}

		req := bridge.BridgeMintReq{
			Sender:     common.HexToAddress(b.senderAddress),
			Receiver:   common.HexToAddress(b.receiverAddress),
			Token:      common.HexToAddress(b.destinationContractAddress),
			Amount:     amountI,
			Fee:        fee,
			RefChainId: big.NewInt(int64(b.sourceNetworkId)),
		}
		copy(req.BurnId[:], common.Hex2BytesFixed(strings.TrimPrefix(b.burnId, "0x"), 32))

		reqs = append(reqs, req)
	}

	signatures, err := client.erc20MultipleMintSigner(client.authenticator.From, client.bridgeContractAddressErc20, big.NewInt(int64(destinationNetworkId)), reqs)
	if err != nil {
		return "", err
	}

	signs := make([][]byte, 0, len(signatures))
	for _, s := range signatures {
		signs = append(signs, hexutil.MustDecode(s))
	}

	ts, fs := make([]common.Address, 0, len(m)), make([]*big.Int, 0, len(m))
	for k, v := range m {
		ts = append(ts, k)
		fs = append(fs, v.Fee)
	}

	return sendRawTransaction(abiObject20, client.rpcClient, destinationNetworkId, "mintTokens", client.authenticator.From, client.bridgeContractAddressErc20, client.authenticator.Signer, reqs, signs, ts, fs)
}

func MintToken(destinationNetworkId, sourceNetworkId networkId, burnId, destinationContractAddress, senderAddress, receiverAddress, amountS, feeS string) (string, error) {
	client, found := nodeClients[destinationNetworkId]
	if !found {
		return "", fmt.Errorf("client not found: %d", destinationNetworkId)
	}

	amount, ok := big.NewInt(0).SetString(amountS, 0)
	if !ok {
		return "", fmt.Errorf("invalid amount %s", amountS)
	}

	fee, ok := big.NewInt(0).SetString(feeS, 0)
	if !ok {
		return "", fmt.Errorf("invalid fee %s", feeS)
	}

	req := bridge.BridgeMintReq{
		Sender:     common.HexToAddress(senderAddress),
		Receiver:   common.HexToAddress(receiverAddress),
		Token:      common.HexToAddress(destinationContractAddress),
		Amount:     amount,
		Fee:        fee,
		RefChainId: big.NewInt(int64(sourceNetworkId)),
	}

	copy(req.BurnId[:], common.Hex2BytesFixed(strings.TrimPrefix(burnId, "0x"), 32))

	signature, err := client.erc20MintSigner(client.authenticator.From, client.bridgeContractAddressErc20, senderAddress, receiverAddress, destinationContractAddress, req.BurnId, req.Amount, req.Fee, req.RefChainId, big.NewInt(int64(destinationNetworkId)))
	if err != nil {
		return "", err
	}

	return sendRawTransaction(abiObject20, client.rpcClient, destinationNetworkId, "mintToken", client.authenticator.From, client.bridgeContractAddressErc20, client.authenticator.Signer, req, hexutil.MustDecode(signature))
}

func sentTo(sourceNetworkId, destinationNetworkId networkId, destinationContractAddress, receiverAddress string, tokenId uint64) (string, error) {
	client, found := nodeClients[destinationNetworkId]
	if !found {
		return "", fmt.Errorf("client not found: %d", destinationNetworkId)
	}

	// edge
	if destinationNetworkId == 100 {
		return sendRawTransaction(abiObject721, client.rpcClient, destinationNetworkId, "sendTo", client.authenticator.From, client.bridgeContractAddressErc721, client.authenticator.Signer, big.NewInt(int64(sourceNetworkId)), common.HexToAddress(destinationContractAddress), big.NewInt(int64(tokenId)), common.HexToAddress(receiverAddress))
	}

	if destinationNetworkId == 280 {
		gasPrice, err := client.rpcClient.SuggestGasPrice(context.Background())
		if err != nil {
			return "", fmt.Errorf("client.rpcClient.SuggestGasPrice error. err:%s", err.Error())
		}
		client.authenticator.GasPrice = gasPrice
		client.authenticator.GasLimit = 0
	}

	transaction, err := client.contractCallerErc721.SendTo(client.authenticator, big.NewInt(int64(sourceNetworkId)), common.HexToAddress(destinationContractAddress), big.NewInt(int64(tokenId)), common.HexToAddress(receiverAddress))
	if err != nil {
		return "", err
	}
	return transaction.Hash().Hex(), nil
}

func refund(sourceNetworkId networkId, sourceContractAddress, senderAddress string, tokenId uint64) (string, error) {
	client, found := nodeClients[sourceNetworkId]
	if !found {
		return "", fmt.Errorf("client not found: %d", sourceNetworkId)
	}

	// edge
	if sourceNetworkId == 100 {
		return sendRawTransaction(abiObject721, client.rpcClient, sourceNetworkId, "bridgeCallBack", client.authenticator.From, client.bridgeContractAddressErc721, client.authenticator.Signer, common.HexToAddress(senderAddress), common.HexToAddress(sourceContractAddress), big.NewInt(int64(tokenId)))
	}

	transaction, err := client.contractCallerErc721.BridgeCallBack(client.authenticator, common.HexToAddress(senderAddress), common.HexToAddress(sourceContractAddress), big.NewInt(int64(tokenId)))
	if err != nil {
		return "", err
	}
	return transaction.Hash().Hex(), nil
}

func getTransactionResult(ctx context.Context, networkId networkId, hash string) (receiptStatus, uint64, error) {
	client, found := nodeClients[networkId]
	if !found {
		return 0, 0, fmt.Errorf("client not found: %d", networkId)
	}
	receipt, err := client.rpcClient.TransactionReceipt(ctx, common.HexToHash(hash))
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return unknown, 0, nil
		}
		return 0, 0, err
	}
	if receipt == nil {
		return 0, 0, fmt.Errorf("getTransactionResult receipt nil. networkID:%d", networkId)
	}
	switch receipt.Status {
	case 0:
		return fail, receipt.BlockNumber.Uint64(), nil
	case 1:
		return success, receipt.BlockNumber.Uint64(), nil
	default:
		return 0, receipt.BlockNumber.Uint64(), fmt.Errorf("invalid receipt status: %d hash: %s network id: %d", receipt.Status, receipt.TxHash.Hex(), networkId)
	}
}

func sendRawTransaction(abi *abi.ABI, client *ethclient.Client, network networkId, method string, sender, contract common.Address, signer bind.SignerFn, params ...interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	input, err := abi.Pack(method, params...)
	if err != nil {
		return "", err
	}

	nonce, err := client.PendingNonceAt(ctx, sender)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}

	baseTx := &types.LegacyTx{
		Nonce:    nonce,
		To:       &contract,
		Data:     input,
		GasPrice: gasPrice,
		Gas:      defaultGas,
		Value:    new(big.Int),
	}

	signedTx, err := signer(sender, types.NewTx(baseTx))
	if err != nil {
		return "", err
	}

	l, f := senderLocker[network]
	if !f || l == nil {
		return "", fmt.Errorf("wierd! Can't find locker of network %d", network)
	}

	locker, found := l[sender]
	if !found || locker == nil {
		return "", fmt.Errorf("wierd! Can't find locker of sender %s", sender.String())
	}

	locker.Lock()
	defer locker.Unlock()

	return signedTx.Hash().Hex(), client.SendTransaction(ctx, signedTx)
}
