package lock

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

const lockerName = "/bridge-backend-service"

type EtcdLocker struct {
	locker *concurrency.Mutex
}

func (el *EtcdLocker) Lock(ctx context.Context) error {
	return el.locker.Lock(ctx)
}

func (el *EtcdLocker) Unlock(ctx context.Context) error {
	return el.locker.Unlock(ctx)
}

func NewEtcdLocker(endpoints []string) (*EtcdLocker, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	session, err := concurrency.NewSession(client)
	if err != nil {
		return nil, err
	}

	locker := concurrency.NewMutex(session, lockerName)

	return &EtcdLocker{locker: locker}, nil
}
