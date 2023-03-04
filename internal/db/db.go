package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/srfbogomolov/warehouse_api/internal/config"
)

func ConnectDB(cfg *config.DB) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(cfg.DRIVER, cfg.DSN)
	if err != nil {
		return
	}

	return
}
