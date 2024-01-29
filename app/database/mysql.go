package database

import (
	"github.com/ApeGame/bridge-backend/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlClient *gorm.DB

func InitializeMysqlDatabase() error {
	client, err := gorm.Open(mysql.Open(config.GetDB().Source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	if err := client.AutoMigrate(
		&Erc721BridgeContractAddress{},
		&Erc20BridgeContractAddress{},
		&BridgeHistory{},
		&SynchronizationProgressRecord{},
		&BridgeHistoryExtra{},
		&SyncZkProgressRecord{},
		&BlackList{},
		&WhiteList{},
	); err != nil {
		return err
	}

	mysqlClient = client

	return nil
}

func GetMysqlClient() *gorm.DB {
	return mysqlClient
}
