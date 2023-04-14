package event

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinebox/graphql"
	"testing"
)

const graphqlEndpoint = `https://testnet.ankr.com/graphql`
const rpcEndpoint = `https://rpc.ankr.com/polygon_mumbai/128bdedab70a53096c6b5132d94384254aee84b8491502b928ab6c08652a7b78`

func TestNewEventFetchThroughRpc(t *testing.T) {
	rpcClient, err := ethclient.Dial("https://eth-goerli-do.ankr.com/rpc")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := NewEventFetchThroughRpc(rpcClient, []string{}, 5, 10, 0, 0, nil); err != nil {
		t.Fatal(err)
	}
}

func TestNewEventFetchThroughGraphQL(t *testing.T) {

	rpcClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	graphClient := graphql.NewClient(graphqlEndpoint)
	graphClient.Log = func(s string) {
		t.Log(s)
	}

	if _, err := NewEventFetchThroughGraphQL(rpcClient, graphClient, []string{}, 5, 10, 0, 0, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFetchThroughGraphQL(t *testing.T) {

	graphClient := graphql.NewClient(graphqlEndpoint)
	graphClient.Log = func(s string) {
		t.Log(s)
	}

	var graphQLResponse eventLogGraphQueryResponseModel

	if err := graphClient.Run(context.Background(), graphql.NewRequest(fmt.Sprintf(getEventLogGraphQueryWithFilters([]string{"0x21a4813D3A13fD762d00FC3551D67553453c5c19", "0xc09d350573715CD441791603c4F01a59Dd832699"}), 5274520, 5274530)), &graphQLResponse); err != nil {
		t.Fatal(err)
	}

	t.Log(graphQLResponse)

	if graphQLResponse.Logs[0].Transaction.Block.Number != 5274525 {
		t.Fail()
	}
}

func TestFetchBlockNumber(t *testing.T) {
	rpcClient, err := ethclient.Dial("http://154.212.140.43:16060")
	if err != nil {
		t.Fatal(err)
	}
	height, err := rpcClient.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(height)
}
