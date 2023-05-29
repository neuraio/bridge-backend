# bridge-backend-service

This project is designed to provide bridge between different blockchain networks

1. Event Driven Design. Watch "NFT Mint" event from the source chain(compare with the block chain of orther side which called "destination") and make sure this NFT token which is going to be minted is not existed in the destination blockchain network. Then call the source bridge contract to allow/discard this mint progress, at the same time,record this NFT token's status so that it will be known as "already minted".
2. Watch "NFT Bridge" event from the source bridge contract which was deployed on the source blockchain network and call another bridge contract which was existed in the blockchain network which is going to bridge to.
3. Provide an API to query each bridge records.
4. Use Mysql as persistence database and etcd for distribute lock

API

```
GET /v1/nft/history-records 获取历史数据
```

```
请求参数
type ListHistoryRecordsFilter struct {
	Address   string            `form:"address"` // 用户地址
	Protocol  database.Protocol `form:"protocol"` // protocol 类型： erc20 erc721
	offset 
	limit 
}
返回参数
type HistoryRecordResp struct {
	ID                  uint64                `json:"id"`
	SrcNetworkID        int                   `json:"srcNetworkID"`
	SrcContractAddress  string                `json:"srcContractAddress"`
	SrcBlockHeight      uint64                `json:"srcBlockHeight"`
	SrcTransactionHash  string                `json:"srcTransactionHash"`
	SrcAddress          string                `json:"srcAddress"`
	DestNetworkID       int                   `json:"destNetworkID"`
	DestContractAddress string                `json:"destContractAddress"`
	DestBlockHeight     uint64                `json:"destBlockHeight"`
	DestTransactionHash string                `json:"destTransactionHash"`
	DestAddress         string                `json:"destAddress"`
	TokenID             uint64                `json:"tokenID"`
	Status              database.BridgeStatus `json:"status"`
	Erc20Amount         string                `json:"erc20Amount"`
	DepositCount        uint64                `json:"depositCount"` // goerli - zkevm 特有字段 
	// extra
	Proof          string `json:"proof"`            // zksync->goerli 调用finalizeWithdrawal方法需要的字段 status = 11 时 正在同步这些数据
	ProofID        string `json:"proofId"`   // zksync->goerli 调用finalizeWithdrawal方法需要的字段 status = 11 时 正在同步这些数据
	L1BatchNumber  string `json:"l1BatchNumber"`   // zksync->goerli 调用finalizeWithdrawal方法需要的字段 status = 11 时 正在同步这些数据
	L1BatchTxIndex string `json:"l1BatchTxIndex"`   // zksync->goerli 调用finalizeWithdrawal方法需要的字段 status = 11 时 正在同步这些数据
	Message        string `json:"message"`   // zksync->goerli 调用finalizeWithdrawal方法需要的字段 status = 11 时 正在同步这些数据

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

bridge status 
0 NftBridgeUndo 
1 NftBridgePending
2 NftBridgeSuccess
3 NftBridgeFail
4 NftBridgeRefundPending
5 NftBridgeRefundSuccess
6 NftBridgeRefundFail
7 NftBridgeZKing              // zkevm <-> goerli 
8 NftBridgeZKDepositSlow      // zkevm <-> goerli
9 NftBridgeDepositing         // zksync deposit
10 NftBridgeWithdrawing        // zksync withdraw
11 NftBridgeFinalizeWithdrawal // zksync finalizeWithdrawal
```

```https://ape-dev.ankr.com/v1/nft/history-records?address=0xD275360894cEb22729DDEb4d6eaD453871f901CA&offset=0&limit=5&protocol=erc20
   {
    "code":200,
    "msg":"success",
    "data":{
        "items":[
            {
                "id":25076,
                "srcNetworkID":12077,
                "srcContractAddress":"0x0000000000000000000000000000000000000000",
                "srcBlockHeight":7091252,
                "srcTransactionHash":"0xcdd2cd199380668689cf26b00c60efebab0acd4abc0589ee83710a0c546623f5",
                "srcAddress":"0xD275360894cEb22729DDEb4d6eaD453871f901CA",
                "destNetworkID":80001,
                "destContractAddress":"0xe4bc4bb0acfdb0c754d1f652599b98d9cdf2afe9",
                "destBlockHeight":34314129,
                "destTransactionHash":"0xba82dbef8d9c4c328807bf1f1377c833ba9dafa37258bbc4aabeccf9da395e07",
                "destAddress":"0xD275360894cEb22729DDEb4d6eaD453871f901CA",
                "tokenID":0,
                "status":2,
                "erc20Amount":"200000000000000000",
                "depositCount":0,
                "proof":"",
                "proofId":"",
                "l1BatchNumber":"",
                "l1BatchTxIndex":"",
                "message":"",
                "createdAt":"2023-04-13T07:50:34.264Z",
                "updatedAt":"2023-04-13T07:50:45.186Z"
            }
        ],
        "recordCount":0, // pending 的数量 状态有database.NftBridgeUndo, database.NftBridgePending,
			database.NftBridgeZKing, database.NftBridgeDepositing, database.NftBridgeWithdrawing
        "total":11
    }
}
```

```

```

```
/v1/nft/ownerOf/:contract_address/:chain_id/:token_id
```

```
/v1/nft/bridge-pair
```

```
/v1/nft/bridge-fee
```

```
/v1/nft/nft-contracts
```
