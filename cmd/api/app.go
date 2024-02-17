package main

import (
	"github.com/aerosystems/subs-service/internal/config"
	HTTPServer "github.com/aerosystems/subs-service/internal/http"
	RPCServer "github.com/aerosystems/subs-service/internal/infrastructure/rpc"
	"github.com/sirupsen/logrus"
)

type App struct {
	log        *logrus.Logger
	cfg        *config.Config
	httpServer *HTTPServer.Server
	rpcServer  *RPCServer.Server
}

func NewApp(
	log *logrus.Logger,
	cfg *config.Config,
	httpServer *HTTPServer.Server,
) *App {
	return &App{
		log:        log,
		cfg:        cfg,
		httpServer: httpServer,
	}
}
