package request

import (
	"github.com/ApeGame/bridge-backend/app/database"
)

type Pager struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

func (f *Pager) Complete() {
	if f.Offset <= 0 {
		f.Offset = 0
	}
	if f.Limit <= 0 {
		f.Limit = 10
	}
	if f.Limit > 20 {
		f.Limit = 20
	}
}

type ListHistoryRecordsFilter struct {
	Address   string            `form:"address"`
	NetworkID int64             `form:"networkID"`
	Protocol  database.Protocol `form:"protocol"`
	Pager
}
