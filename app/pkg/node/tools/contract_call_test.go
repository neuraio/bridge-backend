package tools

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
	"testing"
)

const rpcEndpoint = "https://bsc-mainnet.web3api.com/v1/KBR2FY9IJ2IXESQMQ45X76BNWDAW2TT3Z3"

func TestOwnerOfCall(t *testing.T) {
	rpc, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	if err := OwnerOfCall(rpc, common.HexToAddress("0xdf7952b35f24acf7fc0487d01c8d5690a60dba07"), 100); err != nil {
		t.Fatal(err)
	}
}

func TestOwnerOfCallPost(t *testing.T) {
	t.Run("token not existed", func(t *testing.T) {
		_, err := Erc721OwnerOfCallPost(rpcEndpoint, "0xdf7952b35f24acf7fc0487d01c8d5690a60dba07", 0)
		if errors.Is(err, TokenNotExistError) {
			t.Log("token existed already")
			return
		} else {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("token existed", func(t *testing.T) {
		owner, err := Erc721OwnerOfCallPost(rpcEndpoint, "0xdf7952b35f24acf7fc0487d01c8d5690a60dba07", 100)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(owner)

		if strings.ToLower(owner) != strings.ToLower("0x17539cCa21C7933Df5c980172d22659B8C345C5A") {
			t.Fail()
		}
	})

	t.Run("ft token balance", func(t *testing.T) {
		balance, err := Erc20BalanceOf(rpcEndpoint, "0xA27e024FA03421d86CD1Dbd48cFf7948B5EcCcbf", "0x734548a9e43d2D564600b1B2ed5bE9C2b911c6aB")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(balance.String())
	})
}

func TestGetProof(t *testing.T) {
	got, err := GetProof("https://testnet.era.zksync.dev", "0x9b6515e714e1700b1190eb5e60b99250dab5a1631c1cde1a371a9b57a8092f53")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(got)
}
