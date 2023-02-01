package lock

import (
	"context"
	"testing"
	"time"
)

func TestEtcdLocker_Lock(t *testing.T) {
	locker, err := NewEtcdLocker([]string{"http://192.168.3.96:2379", "http://192.168.3.96:4001"})
	if err != nil {
		t.Fatal(err)
	}

	exception := []string{"lock 1", "unlock 1", "lock 2"}
	results := make([]string, 0)

	if err := locker.Lock(context.Background()); err != nil {
		t.Fatal(err)
	}
	t.Log("locker 1")
	results = append(results, "lock 1")

	down := make(chan struct{}, 1)

	go func() {
		t.Log("block...")
		locker2, err := NewEtcdLocker([]string{"http://192.168.3.96:2379", "http://192.168.3.96:4001"})
		if err != nil {
			t.Error(err)
			return
		}

		if err := locker2.Lock(context.Background()); err != nil {
			t.Error(err)
			return
		}
		t.Log("locker 2")
		results = append(results, "lock 2")
		down <- struct{}{}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		if err := locker.Unlock(context.Background()); err != nil {
			t.Fatal(err)
		}
		t.Log("un locker 1")
		results = append(results, "unlock 1")
	}()

	<-down

	for i := range results {
		if results[i] != exception[i] {
			t.FailNow()
		}
	}
}
