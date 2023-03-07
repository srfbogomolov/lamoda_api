package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/srfbogomolov/warehouse_api/internal/sqlstore"
)

type SqlPlacementRepository struct {
	db *sqlx.DB
}

func NewSqlPlacementRepository(db *sqlx.DB) *SqlPlacementRepository {
	return &SqlPlacementRepository{db: db}
}

func (r *SqlPlacementRepository) getDB() *sqlx.DB {
	return r.db
}

func (r *SqlPlacementRepository) Save(ctx context.Context, p *models.Placement) (int, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return 0, err
	}
	return sqlstore.SavePlacement(ctx, db, p)
}

func (r *SqlPlacementRepository) FindById(ctx context.Context, id int) (*models.Placement, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return nil, err
	}
	return sqlstore.FindPlacementById(ctx, db, id)
}

func (r *SqlPlacementRepository) Find(ctx context.Context) ([]*models.Placement, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return nil, err
	}
	return sqlstore.FindPlacements(ctx, db)
}

func (r *SqlPlacementRepository) FindQTYSumByProductCode(ctx context.Context, productCode string) (int, error) {
	db, err := getSqlxDatabase(ctx, r)
	if err != nil {
		return 0, err
	}
	return sqlstore.FindPlacementQTYSumByProductCode(ctx, db, productCode)
}

func (r *SqlPlacementRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return inSqlTransaction(ctx, r, fn)
}
