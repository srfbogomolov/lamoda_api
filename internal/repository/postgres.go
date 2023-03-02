package repository

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/config"
)

const DRIVER_NAME = "pgx"

type postgresRepository struct {
	Repository
	db *sqlx.DB
}

func NewPostgresRepository(cfg *config.PostgresConfig) (Repository, error) {
	db, err := sqlx.Connect(DRIVER_NAME, cfg.URL)
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db: db}, nil
}
