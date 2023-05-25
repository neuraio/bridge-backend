package store

import (
	"github.com/ApeGame/bridge-backend/app/database"
	"github.com/ApeGame/bridge-backend/app/pkg/request"
)

func ListHistoryRecords(f *request.ListHistoryRecordsFilter) (data []*database.BridgeHistory2, total int64, err error) {
	data = make([]*database.BridgeHistory2, 0)
	q := database.GetMysqlClient().Model(&database.BridgeHistory2{}).
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
	if err := database.GetMysqlClient().Model(&database.BridgeHistory2{}).
		Where("source_address = ? or destination_address = ?", f.Address, f.Address).
		Where("status = ? or status = ? or status = ?", database.NftBridgeUndo, database.NftBridgePending, database.NftBridgeZKing).
		Where("protocol_type = ?", f.Protocol).
		Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
