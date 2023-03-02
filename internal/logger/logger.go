package logger

import (
	"fmt"

	"github.com/srfbogomolov/warehouse_api/internal/config"
	"go.uber.org/zap"
)

func CreateLogger(cfg *config.Config) (logger *zap.Logger, err error) {
	if cfg.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, fmt.Errorf("error building logger: %w", err)
	}

	return logger, nil
}
