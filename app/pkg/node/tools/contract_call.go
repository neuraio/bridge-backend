package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"io"
	"math/big"
	"net/http"
	"strings"
)

const erc721OwnerOfRequestBody = `{"jsonrpc":"2.0","id":7,"method":"eth_call","params":[{"from":"0x0000000000000000000000000000000000000000","data":"0x6352211e%064x","to":"%s"},"latest"]}`
const erc20BalanceRequestBody = `{"jsonrpc":"2.0","id":7,"method":"eth_call","params":[{"from":"0x0000000000000000000000000000000000000000","data":"0x70a08231000000000000000000000000%s","to":"%s"},"latest"]}`
const balanceRequestBody = `{"jsonrpc":"2.0","id":7,"method":"eth_getBalance","params":["%s","latest"]}`

var TokenNotExistError = errors.New("nonexistent token")

func OwnerOfCall(rpc *ethclient.Client, contractAddress common.Address, _ uint64) error {
	res, err := rpc.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: nil, // todo ABI-encoded
	}, nil)

	if err != nil {
		return err
	}

	fmt.Println(string(res))
	return nil
}

func Erc721OwnerOfCallPost(rpcEndpoint, contractAddress string, tokenId uint64) (string, error) {
	logrus.Debugf("Erc721OwnerOfCallPost contractAddress: %s,tokenId: %d", contractAddress, tokenId)
	res, err := http.Post(rpcEndpoint, "application/json", strings.NewReader(fmt.Sprintf(erc721OwnerOfRequestBody, tokenId, contractAddress)))
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	logrus.Debugf("Erc721OwnerOfCallPost response: %s", string(data))

	if err != nil {
		return "", err
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(data, &result); err != nil {
		logrus.Errorf("invalid response %s,request: %s error: %s", string(data), fmt.Sprintf(erc721OwnerOfRequestBody, tokenId, contractAddress), err.Error())
		return "", err
	}

	r, f := result["result"]
	if f {
		owner, ok := r.(string)
		if !ok || len(owner) != 66 || !strings.HasPrefix(owner, "0x") {
			return "", fmt.Errorf("invalid owner address responsed %v", r)
		}
		return "0x" + owner[26:], nil
	}

	e, f := result["error"]
	if f {
		err, ok := e.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("invalid error message responsed %v", e)
		}

		m, found := err["message"]
		if !found {
			return "", fmt.Errorf("invalid error message responsed %v", e)
		}

		message, ok := m.(string)
		if !ok {
			return "", fmt.Errorf("invalid error message responsed %v", m)
		}

		if strings.Contains(message, "reverted") || strings.Contains(message, "execution") {
			return "", TokenNotExistError
		}

		return "", fmt.Errorf("%v", e)
	}

	return "", fmt.Errorf("invalid response body %v", result)
}

func Erc20BalanceOf(rpcEndpoint, accountAddress, contractAddress string) (*big.Int, error) {
	logrus.Debugf("Erc20BalanceOf contract: %s,account: %s", contractAddress, accountAddress)

	accountAddress = strings.ToLower(accountAddress)
	contractAddress = strings.ToLower(contractAddress)
	q := fmt.Sprintf(balanceRequestBody, accountAddress)
	if contractAddress != "0x0000000000000000000000000000000000000000" {
		q = fmt.Sprintf(erc20BalanceRequestBody, strings.TrimPrefix(accountAddress, "0x"), contractAddress)
	}
	res, err := http.Post(rpcEndpoint, "application/json", strings.NewReader(q))
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	logrus.Debugf("Erc20BalanceOfCallPost request: %s response: %s", q, string(data))

	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(data, &result); err != nil {
		logrus.Errorf("invalid response %s,request: %s error: %s", string(data), q, err.Error())
		return nil, err
	}

	r, f := result["result"]
	if f {
		balance, ok := r.(string)
		if !ok {
			return nil, fmt.Errorf("invalid balance amount responsed %v", r)
		}
		b, ok := new(big.Int).SetString(balance, 0)
		if !ok {
			return nil, fmt.Errorf("invalid balance %s %v", balance, r)
		}
		return b, nil
	}

	e, f := result["error"]
	if f {
		err, ok := e.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid error message responsed %v", e)
		}

		m, found := err["message"]
		if !found {
			return nil, fmt.Errorf("invalid error message responsed %v", e)
		}

		message, ok := m.(string)
		if !ok {
			return nil, fmt.Errorf("invalid error message responsed %v", m)
		}

		return nil, fmt.Errorf("%s", message)
	}

	return nil, fmt.Errorf("invalid response body %v", result)
}
