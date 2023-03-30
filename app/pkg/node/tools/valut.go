package tools

import (
	"bytes"
	"context"
	"fmt"
	bridge "github.com/ApeGame/bridge-backend/app/pkg/node/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/sirupsen/logrus"
)

type VaultSigner struct {
	vaultClient *vault.Client
	chainID     int
}

type SignMintErc20 struct {
	Sender          string
	Receiver        string
	Token           string
	BurnId          string
	Amount          *big.Int
	Fee             *big.Int
	SourceNetworkId *big.Int
}

var once = new(sync.Once)
var client *vault.Client

func NewVaultSigner(endpoint, role, secret string, chainId int) (*VaultSigner, error) {
	once.Do(func() {
		config := vault.DefaultConfig() // modify for more granular configuration
		config.Address = endpoint

		var err error
		client, err = vault.NewClient(config)
		if err != nil {
			panic(err)
		}

		roleID := role
		secretID := &auth.SecretID{FromString: secret}

		_, err = login(client, roleID, secretID)
		if err != nil {
			panic(err)
		}

		go renewToken(client, roleID, secretID)
	})

	return &VaultSigner{vaultClient: client, chainID: chainId}, nil
}

func (vs *VaultSigner) Signer(address common.Address, tx *types.Transaction) (*types.Transaction, error) {

	txSerialization, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}
	logrus.Debug("transaction type: ", tx.Type())
	res, err := vs.vaultClient.Logical().Write("/ethereum/tx-sign/private-key", map[string]interface{}{
		"chainID":        fmt.Sprint(vs.chainID),
		"tx":             string(txSerialization),
		"privateKeyName": address.Hex(),
	})
	if err != nil {
		logrus.Errorf("sign transaction through valut error. chain id: %d,sender address: %s,request: %s", vs.chainID, address.Hex(), string(txSerialization))
		return nil, err
	}

	logrus.Debugf("sign transaction. request transaction: %s, response data: %v", string(txSerialization), res.Data)

	signed, found := res.Data["signedTransaction"]
	if !found {
		return nil, fmt.Errorf("not found signed transaction. raw response from valut: %v", res.Data)
	}

	signedTx, ok := signed.(string)
	if !ok {
		return nil, fmt.Errorf("signed tx responsed from valut is invalid. raw response from valut: %v", res.Data)
	}

	return getSignTxByVault(signedTx)
}

func (vs *VaultSigner) SignMintErc20(address, bridge common.Address, sender, receiver, token string, burnId [32]byte, amount, fee, sourceNetworkId, destinationNetworkId *big.Int) (string, error) {
	hash := crypto.Keccak256Hash(
		common.HexToAddress(sender).Bytes(),
		common.HexToAddress(receiver).Bytes(),
		common.HexToAddress(token).Bytes(),
		common.LeftPadBytes(amount.Bytes(), 32),
		common.LeftPadBytes(fee.Bytes(), 32),
		common.LeftPadBytes(sourceNetworkId.Bytes(), 32),
		burnId[:],
		common.LeftPadBytes(destinationNetworkId.Bytes(), 32),
		bridge.Bytes(),
	)

	res, err := vs.vaultClient.Logical().Write("/ethereum/withdraw/byte/sign", map[string]interface{}{
		"privateKeyName": address.Hex(),
		"digestHash":     hash.String(),
	})
	if err != nil {
		return "", err
	}

	signed, found := res.Data["sign"]
	if !found {
		return "", fmt.Errorf("not found signed transaction. raw response from valut: %v", res.Data)
	}

	signedTx, ok := signed.(string)
	if !ok {
		return "", fmt.Errorf("signed tx responsed from valut is invalid. raw response from valut: %v", res.Data)
	}

	return signedTx, nil
}

func (vs *VaultSigner) SignMintErc20S(address, bridgeContract common.Address, destinationNetworkId *big.Int, bridges []bridge.BridgeMintReq) ([]string, error) {
	hashes := make([]string, 0, len(bridges))
	for _, b := range bridges {
		hash, err := vs.SignMintErc20(address, bridgeContract, b.Sender.Hex(), b.Receiver.Hex(), b.Token.Hex(), b.BurnId, b.Amount, b.Fee, b.RefChainId, destinationNetworkId)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, hash)
	}
	return hashes, nil
}

func getSignTxByVault(txSign string) (*types.Transaction, error) {
	txB, err := hexutil.Decode(txSign)
	signedTx := new(types.Transaction)
	rlpStream := rlp.NewStream(bytes.NewBuffer(txB), 0)
	if err = signedTx.DecodeRLP(rlpStream); err != nil {
		return nil, err
	}
	return signedTx, nil
}

func renewToken(client *vault.Client, roleID string, secretID *auth.SecretID) {
	for true {
		vaultLoginResp, err := login(client, roleID, secretID)
		if err != nil {
			logrus.Fatalf("unable to authenticate to Vault: %v", err)
		}
		tokenErr := manageTokenLifecycle(client, vaultLoginResp)
		if tokenErr != nil {
			logrus.Fatalf("unable to start managing token lifecycle: %v", tokenErr)
		}
	}
}

// Starts token lifecycle management. Returns only fatal errors as errors,
// otherwise returns nil, so we can attempt login again.
func manageTokenLifecycle(client *vault.Client, token *vault.Secret) error {
	renew := token.Auth.Renewable // You may notice a different top-level field called Renewable. That one is used for dynamic secrets renewal, not token renewal.
	if !renew {
		logrus.Error("Token is not configured to be renewable. Re-attempting login.")
		return nil
	}

	watcher, err := client.NewLifetimeWatcher(&vault.LifetimeWatcherInput{
		Secret: token,
		// Increment: 86400, // Learn more about this optional value in https://www.vaultproject.io/docs/concepts/lease#lease-durations-and-renewal
	})
	if err != nil {
		return fmt.Errorf("unable to initialize new lifetime watcher for renewing auth token: %w", err)
	}

	go watcher.Start()
	defer watcher.Stop()

	for {
		select {
		// `DoneCh` will return if renewal fails, or if the remaining lease
		// duration is under a built-in threshold and either renewing is not
		// extending it or renewing is disabled. In any case, the caller
		// needs to attempt to log in again.
		case err := <-watcher.DoneCh():
			if err != nil {
				logrus.Errorf("Failed to renew token: %v. Re-attempting login.", err)
				return nil
			}
			// This occurs once the token has reached max TTL.
			logrus.Error("Token can no longer be renewed. Re-attempting login.")
			return nil

		// Successfully completed renewal
		case renewal := <-watcher.RenewCh():
			logrus.Infof("Successfully renewed: %#v", renewal)
		}
	}
}

func login(client *vault.Client, roleID string, secretID *auth.SecretID) (*vault.Secret, error) {
	appRoleAuth, err := auth.NewAppRoleAuth(
		roleID,
		secretID,
	)

	authInfo, err := client.Auth().Login(context.TODO(), appRoleAuth)
	if err != nil {
		return nil, fmt.Errorf("unable to login to userpass auth method: %w", err)
	}
	if authInfo == nil {
		return nil, fmt.Errorf("no auth info was returned after login")
	}

	return authInfo, nil
}
