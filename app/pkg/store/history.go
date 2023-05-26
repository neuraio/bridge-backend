package store

import (
	"github.com/ApeGame/bridge-backend/app/database"
	"github.com/ApeGame/bridge-backend/app/pkg/request"
)

func ListHistoryRecords(f *request.ListHistoryRecordsFilter) (data []*database.BridgeHistory, total int64, err error) {
	data = make([]*database.BridgeHistory, 0)
	q := database.GetMysqlClient().Model(&database.BridgeHistory{}).
		Where("source_address = ? or destination_address = ?", f.Address, f.Address).
		Where("protocol_type = ?", f.Protocol)
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err = q.Offset(f.Offset).Limit(f.Limit).Order("created_at desc").Find(&data).Error; err != nil {
		return nil, 0, err

	}
	return data, total, nil
}

func CountPendingRecord(f *request.ListHistoryRecordsFilter) (int64, error) {
	var total int64
	if err := database.GetMysqlClient().Model(&database.BridgeHistory{}).
		Where("source_address = ? or destination_address = ?", f.Address, f.Address).
		Where("status in ?", []database.BridgeStatus{database.NftBridgeUndo, database.NftBridgePending,
			database.NftBridgeZKing, database.NftBridgeDepositing, database.NftBridgeWithdrawing}).
		Where("protocol_type = ?", f.Protocol).
		Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func GetBridgeHistoryExtras(ids []uint) ([]*database.BridgeHistoryExtra, error) {
	data := make([]*database.BridgeHistoryExtra, 0)
	if err := database.GetMysqlClient().Model(&database.BridgeHistoryExtra{}).
		Where("id in (?)", ids).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
