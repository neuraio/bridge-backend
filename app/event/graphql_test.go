package event

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinebox/graphql"
	"testing"
)

const graphqlEndpoint = `https://testnet.ankr.com/graphql`
const rpcEndpoint = `https://testnet.ankr.com`

func TestNewEventFetchThroughGraphQL(t *testing.T) {

	rpcClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	graphClient := graphql.NewClient(graphqlEndpoint)
	graphClient.Log = func(s string) {
		t.Log(s)
	}

	if _, err := NewEventFetchThroughGraphQL(rpcClient, graphClient, 5, 10, 0, 0, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFetchThroughGraphQL(t *testing.T) {

	graphClient := graphql.NewClient(graphqlEndpoint)
	graphClient.Log = func(s string) {
		t.Log(s)
	}

	var graphQLResponse eventLogGraphQueryResponseModel

	if err := graphClient.Run(context.Background(), graphql.NewRequest(fmt.Sprintf(eventLogGraphQuery, 1088056, 1088060)), &graphQLResponse); err != nil {
		t.Fatal(err)
	}

	t.Log(graphQLResponse)

	if graphQLResponse.Blocks[0].Number.(float64) != 1088056 {
		t.FailNow()
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
