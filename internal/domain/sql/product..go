package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/srfbogomolov/warehouse_api/internal/domain/sql/sqlstore"
)

type SqlProductRepository struct {
	db *sqlx.DB
}

func NewSqlProductRepository(db *sqlx.DB) *SqlProductRepository {
	return &SqlProductRepository{db: db}
}

func (r *SqlProductRepository) getDB() *sqlx.DB {
	return r.db
}

func (r *SqlProductRepository) Save(ctx context.Context, product domain.Product) error {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return err
	}
	productDto := sqlstore.FromProduct(product)
	return sqlstore.SaveProduct(ctx, db, productDto)
}

func (r *SqlProductRepository) GetTotalByItemId(ctx context.Context, id uint) (domain.Product, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return domain.Product{}, err
	}
	return sqlstore.GetProductTotalByItemId(ctx, db, id)
}

func (r *SqlProductRepository) GetTotalByItemCode(ctx context.Context, code uuid.UUID) (domain.Product, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return domain.Product{}, err
	}
	return sqlstore.GetProductTotalByItemCode(ctx, db, code)
}

func (r *SqlProductRepository) GetInStockByItemId(ctx context.Context, id uint, warehouse_id uint) (domain.Product, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return domain.Product{}, err
	}
	return sqlstore.GetProductInStockByItemId(ctx, db, id, warehouse_id)
}

func (r *SqlProductRepository) GetInStockByItemCode(ctx context.Context, code uuid.UUID, warehouse_id uint) (domain.Product, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return domain.Product{}, err
	}
	return sqlstore.GetProductInStockByItemCode(ctx, db, code, warehouse_id)
}

func (r *SqlProductRepository) GetAllInStockByWarehouseId(ctx context.Context, warehouse_id uint) (domain.Products, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return domain.Products{}, err
	}
	return sqlstore.GetAllProductsInStockByWarehouseId(ctx, db, warehouse_id)
}

func (r *SqlProductRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return inSqlTransaction(ctx, r, fn)
}
