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
	"io/ioutil"
	"net/http"
	"strings"
)

const ownerOfHexCode = "0x6352211e%064x"
const erc721OwnerOfRequestBody = `{"jsonrpc":"2.0","id":7,"method":"eth_call","params":[{"from":"0x0000000000000000000000000000000000000000","data":"0x6352211e%064x","to":"%s"},"latest"]}`

var TokenNotExistError = errors.New("nonexistent token")

func OwnerOfCall(rpc *ethclient.Client, contractAddress common.Address, tokenId uint64) error {

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
	data, err := ioutil.ReadAll(res.Body)
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
