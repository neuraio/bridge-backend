package event

import (
	bridge "github.com/ApeGame/bridge-backend/app/pkg/node/abi"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"testing"
)

const destinationNetworkId = 80001
const bridgeAddressNft = "0xc09d350573715CD441791603c4F01a59Dd832699"
const bridgeAddressFt = "0x083DDdD02835385A239b710648f33781B15Cc3C0"
const adminAddress = "0xdab136D1AAceF7417D32ae6B8b13651dbA4Dd580"

func mockNodes(t *testing.T) {
	rpcClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	signer, err := tools.NewVaultSigner("https://vault-dev.ankr.com", "41524d66-95b0-9523-136e-fd24ee9fcac5", "ae5ee054-868d-c931-5c41-0d13d6229ebf", destinationNetworkId)
	if err != nil {
		t.Fatal(err)
	}

	opt := tools.GenerateTransactorOpts(common.HexToAddress(adminAddress), signer.Signer)

	opt.GasLimit = defaultGas
	opt.Signer = signer.Signer

	caller721, err := bridge.NewBridgeTransactor(common.HexToAddress(bridgeAddressNft), rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	caller20, err := bridge.NewBridgeerc20Transactor(common.HexToAddress(bridgeAddressFt), rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	nodeClients[destinationNetworkId] = &contractCaller{
		rpcEndpoint:                 rpcEndpoint,
		rpcClient:                   rpcClient,
		authenticator:               opt,
		contractCallerErc721:        caller721,
		contractCallerErc20:         caller20,
		bridgeContractAddressErc721: common.HexToAddress(bridgeAddressNft),
		bridgeContractAddressErc20:  common.HexToAddress(bridgeAddressFt),
		erc20MintSigner:             signer.SignMintErc20,
		erc20MultipleMintSigner:     signer.SignMintErc20S,
	}

	senderLocker[destinationNetworkId] = make(map[common.Address]*sync.Mutex)
	senderLocker[destinationNetworkId][common.HexToAddress(adminAddress)] = new(sync.Mutex)
}

func TestMintToken(t *testing.T) {
	mockNodes(t)
	hash, err := MintToken(97, 12077, "0xfadb132a33d41e02ace40967f44e5e52846fe341b84d3bddfc072e1df391abcd",
		"0xf34f4756859a8ce2d55bb48d18ef6e470e02787c", "0xb5Dd9E01F02554a6C3d397d0Fd878043e9fa9574",
		"0xb5Dd9E01F02554a6C3d397d0Fd878043e9fa9574", "10000000000000000000", "0")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success hash: ", hash)
}

func TestMintTokens(t *testing.T) {
	mockNodes(t)
	bridges := make([]*ercBridge, 0)
	bridges = append(bridges, &ercBridge{
		sourceNetworkId:            12077,
		burnId:                     "0x8607914fb244e6dfc4f81237cd9b4d82d2f02cb352e78609b7ec9ed6bba704b7",
		destinationContractAddress: "0x82f708bffb087681b81a69b9b665ea863448829f",
		senderAddress:              "0xb5Dd9E01F02554a6C3d397d0Fd878043e9fa9574",
		receiverAddress:            "0xb5Dd9E01F02554a6C3d397d0Fd878043e9fa9574",
		amountS:                    "300000000000000000000",
		feeS:                       "600000000000000000",
	})
	hash, err := MintTokens(destinationNetworkId, bridges)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success hash: ", hash)
}
