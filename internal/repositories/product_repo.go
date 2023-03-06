package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/srfbogomolov/warehouse_api/internal/sqlstore"
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

func (r *SqlProductRepository) Save(ctx context.Context, p *models.Product) error {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return err
	}
	return sqlstore.SaveProduct(ctx, db, p)
}

func (r *SqlProductRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return nil, err
	}
	return sqlstore.GetProductByID(ctx, db, id)
}

func (r *SqlProductRepository) GetAll(ctx context.Context) ([]*models.Product, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return nil, err
	}
	return sqlstore.GetAllProducts(ctx, db)
}

func (r *SqlProductRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return inSqlTransaction(ctx, r, fn)
}
