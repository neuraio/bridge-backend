package pkg

import (
	"context"
	"github.com/ApeGame/bridge-backend/app/config"
	"github.com/ApeGame/bridge-backend/app/pkg/api"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Init()

	formatter := logrus.JSONFormatter{
		PrettyPrint: false,
	}
	logrus.SetFormatter(&formatter)
	logrus.SetLevel(logrus.DebugLevel)
}

type Server struct {
	engine *gin.Engine
}

func NewServer(endpoints []string) *Server {

	for i := range endpoints {
		client, err := ethclient.Dial(endpoints[i])
		if err != nil {
			logrus.Fatalln(err)
		}

		chainId, err := client.ChainID(context.Background())
		if err != nil {
			logrus.Fatalln(err)
		}

		logrus.Infof("block chains. endpoint: %s, chain id: %d", endpoints[i], chainId.Int64())

		ethClient := &api.EthRpcClient{
			Endpoint: endpoints[i],
			ChainId:  chainId.Int64(),
			Client:   client,
		}

		api.EthRpc[chainId.Int64()] = ethClient
	}

	e := gin.Default()
	api.SetupRouter(e)

	return &Server{
		engine: e,
	}
}

func (s *Server) Start() {
	logrus.Infof("start to listen %s", config.Listen())
	_ = s.engine.Run(config.Listen())
}
