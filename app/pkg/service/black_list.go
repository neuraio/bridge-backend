package service

import (
	"fmt"
	"github.com/ApeGame/bridge-backend/app/database"
	"gorm.io/gorm"
	"strings"
	"sync"
)

type Blacklist struct {
	locker    sync.Locker
	db        *gorm.DB
	addresses []string
}

var BlackList *Blacklist

func InitBlackList(db *gorm.DB) error {
	list := make([]*database.BlackList, 0)
	if err := db.Find(&list).Error; err != nil {
		return err
	}
	BlackList = &Blacklist{
		locker:    new(sync.Mutex),
		db:        db,
		addresses: make([]string, 0, len(list)),
	}
	for i := range list {
		BlackList.addresses = append(BlackList.addresses, strings.ToLower(list[i].CryptoAddress))
	}

	return nil
}

func (bl *Blacklist) Add(address string) error {
	bl.locker.Lock()
	defer bl.locker.Unlock()

	if bl.Check(address) {
		return fmt.Errorf("address %s is already in black list", address)
	}
	bl.addresses = append(bl.addresses, strings.ToLower(address))
	return bl.db.Save(&database.BlackList{
		CryptoAddress: address,
	}).Error
}

func (bl *Blacklist) Check(address string) bool {
	for i := range bl.addresses {
		if bl.addresses[i] == strings.ToLower(address) {
			return true
		}
	}
	return false
}
