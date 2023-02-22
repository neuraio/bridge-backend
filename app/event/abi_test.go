package event

import (
	bridge "github.com/ApeGame/bridge-backend/app/pkg/node/abi"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"testing"
)

const destinationNetworkId = 12077
const bridgeAddressNft = "0xc09d350573715CD441791603c4F01a59Dd832699"
const bridgeAddressFt = "0x21a4813D3A13fD762d00FC3551D67553453c5c19"
const adminAddress = "0xdab136D1AAceF7417D32ae6B8b13651dbA4Dd580"

func mockNodes(t *testing.T) {
	rpcClient, err := ethclient.Dial("https://testnet.ankr.com")
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
		sourceNetworkId:            97,
		burnId:                     "0x8250c44bb4b5c670e16d696bf706c55df519688a8ab11eb6ea0778db7b729a33",
		destinationContractAddress: "0x05ecb110a232161fbf719e5889a95e42ca0be154",
		senderAddress:              "0x5257ad0c2dCFfCC0B3a30eAed151E76fF441Eb60",
		receiverAddress:            "0x5257ad0c2dCFfCC0B3a30eAed151E76fF441Eb60",
		amountS:                    "3000000000000000000",
		feeS:                       "6000000000000000",
	})
	hash, err := MintTokens(destinationNetworkId, bridges)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success hash: ", hash)
}
