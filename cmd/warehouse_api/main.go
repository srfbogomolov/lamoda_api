package main

import (
	"log"

	"github.com/srfbogomolov/warehouse_api/internal/app"
	"github.com/srfbogomolov/warehouse_api/internal/config"
	"github.com/srfbogomolov/warehouse_api/internal/db"
	"github.com/srfbogomolov/warehouse_api/internal/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	unsugared, err := logger.CreateLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger := unsugared.Sugar()

	db, err := db.ConnectDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg, db, logger)
}
