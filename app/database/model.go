package database

import (
	"gorm.io/gorm"
)

type BridgeStatus int

const (
	NftBridgeUndo BridgeStatus = iota
	NftBridgePending
	NftBridgeSuccess
	NftBridgeFail
	NftBridgeRefundPending
	NftBridgeRefundSuccess
	NftBridgeRefundFail
	NftBridgeZKing
	NftBridgeZKSuccess // 前端成功
	NftBridgeZKDepositSlow
)

type Protocol string

const (
	Erc20  Protocol = "erc20"
	Erc721 Protocol = "erc721"
)

type BridgeHistory struct {
	gorm.Model
	SourceNetworkId            int    `json:"sourceNetworkId"`
	SourceContractAddress      string `json:"sourceContractAddress"`
	SourceBlockHeight          uint64 `json:"sourceBlockHeight"`
	SourceTransactionHash      string `json:"sourceTransactionHash" gorm:"index"`
	SourceAddress              string `json:"sourceAddress" gorm:"index"`
	DestinationNetworkId       int    `json:"destinationNetworkId"`
	DestinationContractAddress string `json:"destinationContractAddress"`
	DestinationBlockHeight     uint64 `json:"destinationBlockHeight"`
	DestinationTransactionHash string `json:"destinationTransactionHash" gorm:"index"`
	DestinationAddress         string `json:"destinationAddress" gorm:"index"`
	FailReason                 string
	RefundHash                 string
	TokenId                    uint64       `json:"tokenId" gorm:"index"`
	Status                     BridgeStatus `json:"status"`
	RefundAudit                string
	ProtocolType               Protocol `json:"protocolType"`
	Erc20BurnId                string   `json:"erc20BurnId"`
	Erc20Amount                string   `json:"erc20Amount"`
	Fee                        string   `json:"fee"`

	DepositCount uint64 // ZK Bridge attribute
}

type Erc721BridgeContractAddress struct {
	gorm.Model
	NetworkId uint64
	Address   string
	Symbol    string
	Name      string
	Icon      string
}

const SynchronizedLatestHeight = "synchronized_latest_height"

type SynchronizationProgressRecord struct {
	gorm.Model
	BlockHeight uint64
	NetworkId   int `gorm:"uniqueIndex"`
	Comment     string
}

type Erc20BridgeContractAddress struct {
	gorm.Model
	Name                  string
	NetworkId             uint64
	ContractAddress       string
	ContractName          string
	ContractIcon          string
	MinBurn               string
	MaxBurn               string
	MinFee                string
	RollupContractAddress string
	DstNetworkId          uint64
}
