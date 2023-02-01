package tools

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Credential struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    common.Address
}

func HexToCredential(hexKey string) (*Credential, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}
	return KeyToCredential(privateKey)
}

func KeyToCredential(privateKey *ecdsa.PrivateKey) (*Credential, error) {
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key cast error")
	}
	address := crypto.PubkeyToAddress(*publicKey)

	credential := &Credential{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
	}
	return credential, nil
}

func GenerateTransactorOpts(sender common.Address, signer func(address common.Address, tx *types.Transaction) (*types.Transaction, error)) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:    sender,
		Signer:  signer,
		Context: context.Background(),
	}
}
