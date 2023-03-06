package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/srfbogomolov/warehouse_api/internal/config"
)

func ConnectDB(cfg *config.DB) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	return
}
