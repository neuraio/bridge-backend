package event

import (
	"errors"
	"github.com/ApeGame/bridge-backend/app/database"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type synchronizeProgressRecord interface {
	startRecord() error
	latestHeight() (uint64, error)
	autoCommitRecord(height uint64) (string, error)
}

type synchronizeProgressMysqlRecorder struct {
	mysqlDatabaseClient dataRecorderTransaction
	dataModel           *database.SynchronizationProgressRecord
}

func (s synchronizeProgressMysqlRecorder) startRecord() error {
	return s.mysqlDatabaseClient.begin()
}

func (s synchronizeProgressMysqlRecorder) latestHeight() (uint64, error) {
	if err := database.GetMysqlClient().Where("id = ?", s.dataModel.ID).First(s.dataModel).Error; err != nil {
		return 0, err
	}
	return s.dataModel.BlockHeight, nil
}

func (s synchronizeProgressMysqlRecorder) autoCommitRecord(height uint64) (string, error) {
	s.dataModel.BlockHeight = height
	rawTransaction := s.mysqlDatabaseClient.getRawClient()
	mysqlClient, ok := rawTransaction.(*gorm.DB)
	if !ok {
		logrus.Fatalln("Weird! convert raw transaction client error")
	}

	if err := mysqlClient.Save(s.dataModel).Error; err != nil {
		logrus.Errorf("update block height %d  %s", height, err.Error())
		s.mysqlDatabaseClient.setError(err)
	}
	return s.mysqlDatabaseClient.autoCommit()
}

func newSynchronizeProgressMysqlRecorder(mysqlTransaction dataRecorderTransaction, networkId networkId) *synchronizeProgressMysqlRecorder {
	var dataModel = new(database.SynchronizationProgressRecord)
	if err := database.GetMysqlClient().Where("network_id = ? and comment = ?", networkId, database.SynchronizedLatestHeight).First(dataModel).Error; err != nil {
		panic(err)
	}
	return &synchronizeProgressMysqlRecorder{
		mysqlDatabaseClient: mysqlTransaction,
		dataModel:           dataModel,
	}
}

type dataRecorderTransaction interface {
	begin() error
	commit() error
	rollback() error
	autoCommit() (string, error)
	getRawClient() interface{}
	setError(err error)
}

type dataRecorderMysqlTransaction struct {
	mysqlClient      *gorm.DB
	mysqlTransaction *gorm.DB
	err              error
}

func newDataRecorderMysqlTransaction(client *gorm.DB) *dataRecorderMysqlTransaction {
	return &dataRecorderMysqlTransaction{mysqlClient: client}
}

func (d *dataRecorderMysqlTransaction) begin() error {
	d.mysqlTransaction = d.mysqlClient.Begin()
	d.err = nil
	return d.mysqlTransaction.Error
}

func (d *dataRecorderMysqlTransaction) setError(err error) {
	d.err = err
}

func (d *dataRecorderMysqlTransaction) commit() error {
	defer func() {
		d.mysqlTransaction = nil
	}()
	if d.mysqlTransaction == nil {
		return errors.New("could not commit a transaction has not been created")
	}
	if err := d.mysqlTransaction.Commit().Error; err == nil {
		return nil
	} else {
		return err
	}
}

func (d *dataRecorderMysqlTransaction) rollback() error {
	defer func() {
		d.mysqlTransaction = nil
	}()
	if d.mysqlTransaction == nil {
		return errors.New("could not rollback a transaction has not been created")
	}
	return d.mysqlTransaction.Rollback().Error
}

func (d *dataRecorderMysqlTransaction) autoCommit() (string, error) {
	defer func() {
		d.mysqlTransaction = nil
	}()
	if d.mysqlTransaction == nil {
		return "", errors.New("could not commit a transaction has not been created")
	}
	if d.err == nil {
		return "", d.mysqlTransaction.Commit().Error
	}
	return d.err.Error(), d.mysqlTransaction.Rollback().Error
}

func (d *dataRecorderMysqlTransaction) getRawClient() interface{} {
	return d.mysqlTransaction
}
