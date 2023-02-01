package service

import (
	"time"

	"github.com/ApeGame/bridge-backend/app/database"
	"github.com/ApeGame/bridge-backend/app/pkg/request"
	"github.com/ApeGame/bridge-backend/app/pkg/store"
)

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

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ListHistoryRecords(f *request.ListHistoryRecordsFilter) (data []*HistoryRecordResp, total, recordCount int64, err error) {
	resp := make([]*HistoryRecordResp, 0)
	if f.Address == "" {
		return resp, 0, 0, nil
	}
	records, total, err := store.ListHistoryRecords(f)
	if err != nil {
		return nil, 0, 0, err
	}
	recordCount, err = store.CountPendingRecord(f)
	if err != nil {
		return nil, 0, 0, err
	}

	for _, record := range records {
		resp = append(resp, &HistoryRecordResp{
			ID:                  uint64(record.ID),
			SrcNetworkID:        record.SourceNetworkId,
			SrcContractAddress:  record.SourceContractAddress,
			SrcBlockHeight:      record.SourceBlockHeight,
			SrcTransactionHash:  record.SourceTransactionHash,
			SrcAddress:          record.SourceAddress,
			DestNetworkID:       record.DestinationNetworkId,
			DestContractAddress: record.DestinationContractAddress,
			DestBlockHeight:     record.DestinationBlockHeight,
			DestTransactionHash: record.DestinationTransactionHash,
			DestAddress:         record.DestinationAddress,
			TokenID:             record.TokenId,
			Status:              record.Status,
			CreatedAt:           record.CreatedAt,
			UpdatedAt:           record.UpdatedAt,
			Erc20Amount:         record.Erc20Amount,
		})
	}

	return resp, total, recordCount, nil
}
