package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/config"
	"github.com/srfbogomolov/warehouse_api/internal/repositories"
	"github.com/srfbogomolov/warehouse_api/internal/services"
	"go.uber.org/zap"
)

const (
	codecContentType = "application/json"
	serviceName      = "warehouse"
	servicePath      = "/warehouse"
)

func Run(cfg *config.Config, db *sqlx.DB, logger *zap.SugaredLogger) {
	warehouseRepo := repositories.NewSqlWarehouseRepository(db)
	productRepo := repositories.NewSqlProductRepository(db)
	placementRepo := repositories.NewSqlPlacementRepository(db)

	service := services.NewService(warehouseRepo, productRepo, placementRepo, logger)
	handler := NewHandler(service)

	logger.Infow("starting server", "port", cfg.Server.Port)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Fatalw("failed to serve", err)
	}
}

func NewHandler(service services.Service) *mux.Router {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), codecContentType)
	server.RegisterService(service, serviceName)

	router := mux.NewRouter()
	router.Handle(servicePath, server)

	return router
}
