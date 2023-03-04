package app

import (
	"github.com/srfbogomolov/warehouse_api/internal/config"
	"github.com/srfbogomolov/warehouse_api/internal/db"
	"go.uber.org/zap"
)

func Run(cfg *config.Config, logger *zap.SugaredLogger) {
	_, err := db.ConnectDB(cfg.DB)
	if err != nil {
		logger.Fatalw("error openning database connection", err)
	}
}
