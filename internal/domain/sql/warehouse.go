package sql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/srfbogomolov/warehouse_api/internal/domain/sql/sqlstore"
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

func (r *SqlWarehouseRepository) Save(ctx context.Context, warehouse domain.Warehouse) error {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return err
	}
	warehouseDto := sqlstore.FromWarehouse(warehouse)

	for _, product := range warehouse.GetProducts() {
		productDto := sqlstore.FromProduct(product)
		sqlstore.SaveProduct(ctx, db, productDto)
		placementId, err := sqlstore.SavePlacement(ctx, db, warehouseDto, productDto)
		if err != nil {
			return err
		}
		sqlstore.SaveReservation(ctx, db, placementId, productDto)
	}

	return sqlstore.SaveWarehouse(ctx, db, warehouseDto)
}

func (r *SqlWarehouseRepository) GetById(ctx context.Context, id uint) (domain.Warehouse, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return domain.Warehouse{}, err
	}
	warehouse, err := sqlstore.GetWarehouseById(ctx, db, id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	products, err := sqlstore.GetAllProductsInStockByWarehouseId(ctx, db, id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	warehouse.SetProducts(products)
	return warehouse, err
}

func (r *SqlWarehouseRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return inSqlTransaction(ctx, r, fn)
}
