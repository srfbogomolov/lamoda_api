package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/srfbogomolov/warehouse_api/internal/domain/sql/sqlstore"
)

type SqlItemRepository struct {
	db *sqlx.DB
}

func NewSqlItemRepository(db *sqlx.DB) *SqlItemRepository {
	return &SqlItemRepository{db: db}
}

func (r *SqlItemRepository) getDB() *sqlx.DB {
	return r.db
}

func (r *SqlItemRepository) GetById(ctx context.Context, id uint) (*domain.Item, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return &domain.Item{}, err
	}
	return sqlstore.GetItemById(ctx, db, id)
}

func (r *SqlItemRepository) GetByCode(ctx context.Context, code uuid.UUID) (*domain.Item, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return &domain.Item{}, err
	}
	return sqlstore.GetItemByCode(ctx, db, code)
}

func (r *SqlItemRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return inSqlTransaction(ctx, r, fn)
}
