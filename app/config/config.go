package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var C config

func Init() {
	viper.AddConfigPath("./etc")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panicf("read config file error: %s", err)
	}
	if err := viper.Unmarshal(&C); err != nil {
		logrus.Panicf("parse config file error: %s", err)
	}
	if C.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Infof("config loaded: %+v", C)
}

func Debug() bool {
	return C.Debug
}

func Listen() string {
	return C.Listen
}

func GetDB() DB {
	return C.DB
}

func GetChainCfg() []ChainCfg {
	return C.ChainCfg
}

func GetEtcdEndpoints() []string {
	return C.EtcdEndpoints
}

func GetVaultConfig() ValutConfig {
	return C.Valut
}
