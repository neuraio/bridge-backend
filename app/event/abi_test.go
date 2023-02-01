package event

import (
	bridge "github.com/ApeGame/bridge-backend/app/pkg/node/abi"
	"github.com/ApeGame/bridge-backend/app/pkg/node/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func mockNodes(t *testing.T) {
	rpcClient, err := ethclient.Dial("http://103.23.44.17:18575")
	if err != nil {
		t.Fatal(err)
	}

	signer, err := tools.NewVaultSigner("https://vault-dev.ankr.com", "41524d66-95b0-9523-136e-fd24ee9fcac2", "78595c1f-08d4-f6df-ab38-4fc97471c0ea", 97)
	if err != nil {
		t.Fatal(err)
	}

	opt := tools.GenerateTransactorOpts(common.HexToAddress("0xd6a5914E2b8676bD8Fd2fcD9c0fD1FCf1B5A9411"), signer.Signer)

	opt.GasLimit = defaultGas
	opt.Signer = signer.Signer

	caller721, err := bridge.NewBridgeTransactor(common.HexToAddress("0x03CBcAa350A7f35E7D60976c9ee9AdC5ed46BF23"), rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	caller20, err := bridge.NewBridgeErc20Transactor(common.HexToAddress("0xe99871132bDC1c8306556f2B23FCE261575d18B3"), rpcClient)
	if err != nil {
		t.Fatal(err)
	}

	nodeClients[97] = &contractCaller{
		rpcClient:                   rpcClient,
		authenticator:               opt,
		contractCallerErc721:        caller721,
		contractCallerErc20:         caller20,
		bridgeContractAddressErc721: common.HexToAddress("0x03CBcAa350A7f35E7D60976c9ee9AdC5ed46BF23"),
		bridgeContractAddressErc20:  common.HexToAddress("0xe99871132bDC1c8306556f2B23FCE261575d18B3"),
		erc20MintSigner:             signer.SignMintErc20,
	}
}

func TestMintToken(t *testing.T) {
	mockNodes(t)
	hash, err := MintToken(97, 12077, "0xfadb132a33d41e02ace40967f44e5e52846fe341b84d3bddfc072e1df391abcd",
		"0xf34f4756859a8ce2d55bb48d18ef6e470e02787c", "0xb5Dd9E01F02554a6C3d397d0Fd878043e9fa9574",
		"0xb5Dd9E01F02554a6C3d397d0Fd878043e9fa9574", "10000000000000000000", 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success hash: ", hash)
}
