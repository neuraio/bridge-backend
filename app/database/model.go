package database

import (
	"gorm.io/gorm"
	"time"
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
	NftBridgeZKDepositSlow
	NftBridgeDepositing         // zksync deposit
	NftBridgeWithdrawing        // zksync withdraw
	NftBridgeFinalizeWithdrawal // zksync finalizeWithdrawal
)

type Protocol string

const (
	Erc20  Protocol = "erc20"
	Erc721 Protocol = "erc721"
	//Erc20Rollup Protocol = "erc20rollup" // eth zksync rollup
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
const SyncZkFinalizeWithdrawalHeight = "sync_zk_event_height"

type SynchronizationProgressRecord struct {
	gorm.Model
	BlockHeight uint64
	NetworkId   int `gorm:"uniqueIndex"`
	Comment     string
}

type SyncZkProgressRecord struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	BlockHeight uint64
	NetworkId   int    `gorm:"primary_key"`
	Type        string `gorm:"primary_key"`
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
	RollupContractAddress string // zkevm polygon eth
	DstNetworkId          uint64 // zkevm polygon eth network id
	LDstNetworkId         uint64 // l1 l2 network id
}

type BridgeHistoryExtra struct {
	ID        uint `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Proof          string
	ProofID        string
	L1BatchNumber  string
	L1BatchTxIndex string
	Message        string
}
