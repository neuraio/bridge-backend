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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

const defaultGas = 400000

type erc20MintSignFn func(address, bridge common.Address, sender, receiver, token, burnId string, amount, sourceNetworkId, destinationNetworkId *big.Int) (string, error)

type contractCaller struct {
	rpcClient                   *ethclient.Client
	contractCallerErc721        *bridge.BridgeTransactor
	contractCallerErc20         *bridge.BridgeErc20Transactor
	authenticator               *bind.TransactOpts
	bridgeContractAddressErc721 common.Address
	bridgeContractAddressErc20  common.Address
	erc20MintSigner             erc20MintSignFn
}

type receiptStatus int

const (
	unknown = iota
	success
	fail
)

var nodeClients = make(map[networkId]*contractCaller)

func registerNodeClients(networkId networkId, bridge721ContractAddress, bridge20ContractAddress string, client *ethclient.Client, address string, signer bind.SignerFn, erc20MintSigner erc20MintSignFn) error {

	opt := tools.GenerateTransactorOpts(common.HexToAddress(address), signer)

	opt.GasLimit = defaultGas
	opt.Signer = signer

	caller721, err := bridge.NewBridgeTransactor(common.HexToAddress(bridge721ContractAddress), client)
	if err != nil {
		return err
	}

	caller20, err := bridge.NewBridgeErc20Transactor(common.HexToAddress(bridge20ContractAddress), client)
	if err != nil {
		return err
	}

	nodeClients[networkId] = &contractCaller{
		rpcClient:                   client,
		authenticator:               opt,
		contractCallerErc721:        caller721,
		contractCallerErc20:         caller20,
		bridgeContractAddressErc721: common.HexToAddress(bridge721ContractAddress),
		bridgeContractAddressErc20:  common.HexToAddress(bridge20ContractAddress),
		erc20MintSigner:             erc20MintSigner,
	}
	return nil
}

func MintToken(destinationNetworkId, sourceNetworkId networkId, burnId, destinationContractAddress, senderAddress, receiverAddress, amountS string, feeRatio float64) (string, error) {
	client, found := nodeClients[destinationNetworkId]
	if !found {
		return "", fmt.Errorf("client not found: %d", destinationNetworkId)
	}

	amount, ok := big.NewInt(0).SetString(amountS, 0)
	if !ok {
		return "", fmt.Errorf("invalid amount %s", amountS)
	}

	fee := big.NewInt(0).Mul(amount, big.NewInt(int64(feeRatio*1e8)))
	fee = fee.Div(fee, big.NewInt(1e8))

	req := bridge.BridgeMintReq{
		Sender:     common.HexToAddress(senderAddress),
		Receiver:   common.HexToAddress(receiverAddress),
		Token:      common.HexToAddress(destinationContractAddress),
		Amount:     amount,
		Fee:        fee,
		RefChainId: big.NewInt(int64(sourceNetworkId)),
	}

	copy(req.BurnId[:], common.Hex2BytesFixed(strings.TrimLeft(burnId, "0x"), 32))

	signature, err := client.erc20MintSigner(client.authenticator.From, client.bridgeContractAddressErc20, senderAddress, receiverAddress, destinationContractAddress, burnId, req.Amount, req.RefChainId, big.NewInt(int64(destinationNetworkId)))
	if err != nil {
		return "", err
	}

	return sendRawTransaction(abiObject20, client.rpcClient, "mintToken", client.authenticator.From, client.bridgeContractAddressErc20, client.authenticator.Signer, req, common.Hex2Bytes(strings.TrimLeft(signature, "0x")))
}

func sentTo(sourceNetworkId, destinationNetworkId networkId, destinationContractAddress, receiverAddress string, tokenId uint64) (string, error) {
	client, found := nodeClients[destinationNetworkId]
	if !found {
		return "", fmt.Errorf("client not found: %d", destinationNetworkId)
	}

	// edge
	if destinationNetworkId == 100 {
		return sendRawTransaction(abiObject721, client.rpcClient, "sendTo", client.authenticator.From, client.bridgeContractAddressErc721, client.authenticator.Signer, big.NewInt(int64(sourceNetworkId)), common.HexToAddress(destinationContractAddress), big.NewInt(int64(tokenId)), common.HexToAddress(receiverAddress))
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
		return sendRawTransaction(abiObject721, client.rpcClient, "bridgeCallBack", client.authenticator.From, client.bridgeContractAddressErc721, client.authenticator.Signer, common.HexToAddress(senderAddress), common.HexToAddress(sourceContractAddress), big.NewInt(int64(tokenId)))
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

	switch receipt.Status {
	case 0:
		return fail, receipt.BlockNumber.Uint64(), nil
	case 1:
		return success, receipt.BlockNumber.Uint64(), nil
	default:
		return 0, receipt.BlockNumber.Uint64(), fmt.Errorf("invalid receipt status: %d hash: %s network id: %d\n", receipt.Status, receipt.TxHash.Hex(), networkId)
	}
}

func sendRawTransaction(abi *abi.ABI, client *ethclient.Client, method string, sender, contract common.Address, signer bind.SignerFn, params ...interface{}) (string, error) {
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

	return signedTx.Hash().Hex(), client.SendTransaction(ctx, signedTx)
}
