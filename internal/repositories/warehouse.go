package repositories

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/srfbogomolov/warehouse_api/internal/sqlstore"

	"github.com/jmoiron/sqlx"
)

type SqlWarehouseRepository struct {
	db *sqlx.DB
}

func NewSqlWarehouseRepository(db *sqlx.DB) *SqlWarehouseRepository {
	return &SqlWarehouseRepository{db: db}
}

func (r *SqlWarehouseRepository) getDB() *sqlx.DB {
	return r.db
}

func (r *SqlWarehouseRepository) Save(ctx context.Context, w *models.Warehouse) error {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return err
	}
	return sqlstore.SaveWarehouse(ctx, db, w)
}

func (r *SqlWarehouseRepository) FindAll(ctx context.Context) ([]*models.Warehouse, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return nil, err
	}
	return sqlstore.FindAllWarehouse(ctx, db)
}

func (r *SqlWarehouseRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return inSqlTransaction(ctx, r, fn)
}
