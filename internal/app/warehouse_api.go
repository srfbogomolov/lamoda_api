package app

import (
	"github.com/srfbogomolov/warehouse_api/internal/config"
	"github.com/srfbogomolov/warehouse_api/internal/repository"
	"go.uber.org/zap"
)

func Run(cfg *config.Config, logger *zap.SugaredLogger) {
	_, err := repository.NewPostgresRepository(cfg.Postgres)
	if err != nil {
		logger.Fatalw("failed to connect to postgres", err)
	}
}
