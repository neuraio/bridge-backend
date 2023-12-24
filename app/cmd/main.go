package main

import (
	"context"
	"github.com/ApeGame/bridge-backend/app/database"
	"github.com/ApeGame/bridge-backend/app/pkg/lock"
	"github.com/ApeGame/bridge-backend/app/pkg/service"
	"os"
	"os/signal"

	"github.com/ApeGame/bridge-backend/app/config"
	"github.com/ApeGame/bridge-backend/app/event"
	"github.com/ApeGame/bridge-backend/app/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := database.InitializeMysqlDatabase(); err != nil {
		logrus.Fatalln(err)
	}

	logrus.Info("database initialize successful")

	if err := service.InitBlackList(database.GetMysqlClient()); err != nil {
		logrus.Fatalln(err)
	}

	logrus.Info("black list initialize successful")

	locker, err := lock.NewEtcdLocker(config.GetEtcdEndpoints())
	if err != nil {
		logrus.Fatalln(err)
	}

	go func() {
		logrus.Infoln("block...")
		if err := locker.Lock(context.Background()); err != nil {
			logrus.Fatalln(err)
		}
		event.StartEventSystem()
		event.StartJobHandler(context.Background())
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Kill, os.Interrupt) // kill & Ctl+C

	endpoints := make([]string, 0, len(config.GetChainCfg()))
	for _, chain := range config.GetChainCfg() {
		endpoints = append(endpoints, chain.RpcUrl)
	}

	go pkg.NewServer(endpoints).Start()

	logrus.Info("graceful shut down: ", <-shutdown)
	if err := locker.Unlock(context.Background()); err != nil {
		logrus.Fatalln(err)
	}
}
