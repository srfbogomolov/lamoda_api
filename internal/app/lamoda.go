package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/srfbogomolov/warehouse_api/internal/config"
	"github.com/srfbogomolov/warehouse_api/internal/services"
	"go.uber.org/zap"
)

const (
	CodecContentType = "application/json"
	ServiceName      = "lamoda"
	ServicePath      = "/lamoda"
)

type App struct {
	service services.Service
	cfg     *config.Config
	logger  *zap.SugaredLogger
}

func NewApp(service services.Service, cfg *config.Config, logger *zap.SugaredLogger) *App {
	return &App{
		service: service,
		cfg:     cfg,
		logger:  logger,
	}
}

func (a *App) Start() {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), CodecContentType)
	server.RegisterService(a.service, ServiceName)

	router := mux.NewRouter()
	router.Handle(ServicePath, server)

	a.logger.Infow("starting server", "port", a.cfg.Server.Port)

	addr := fmt.Sprintf("%s:%s", a.cfg.Server.Host, a.cfg.Server.Port)
	if err := http.ListenAndServe(addr, router); err != nil {
		a.logger.Fatalw("failed to serve", err)
	}
}
