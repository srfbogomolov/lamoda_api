package repositories

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/srfbogomolov/warehouse_api/internal/sqlstore"

	"github.com/jmoiron/sqlx"
)

type sqlWarehouseRepository struct {
	db *sqlx.DB
}

func NewSqlWarehouseRepository(db *sqlx.DB) *sqlWarehouseRepository {
	return &sqlWarehouseRepository{db: db}
}

func (r *sqlWarehouseRepository) getDB() *sqlx.DB {
	return r.db
}

func (r *sqlWarehouseRepository) Save(ctx context.Context, w *models.Warehouse) error {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return err
	}
	return sqlstore.SaveWarehouse(ctx, db, w)
}

func (r *sqlWarehouseRepository) GetAll(ctx context.Context) ([]*models.Warehouse, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return nil, err
	}
	return sqlstore.GetAllWarehouse(ctx, db)
}

func (r *sqlWarehouseRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return inSqlTransaction(ctx, r, fn)
}
