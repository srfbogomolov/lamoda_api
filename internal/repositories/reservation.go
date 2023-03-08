package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/srfbogomolov/warehouse_api/internal/sqlstore"
)

type SqlReservationRepository struct {
	db *sqlx.DB
}

func NewSqlReservationRepository(db *sqlx.DB) *SqlReservationRepository {
	return &SqlReservationRepository{db: db}
}

func (repo *SqlReservationRepository) getDB() *sqlx.DB {
	return repo.db
}

func (repo *SqlReservationRepository) Save(ctx context.Context, reservation *models.Reservation) (uint, error) {
	db, err := getSqlxDatabase(ctx, repo)
	if err != nil {
		return 0, err
	}
	return sqlstore.SaveReservation(ctx, db, reservation)
}

func (repo *SqlReservationRepository) Find(ctx context.Context) ([]*models.Reservation, error) {
	db, err := getSqlxDatabase(ctx, repo)
	if err != nil {
		return nil, err
	}
	return sqlstore.FindReservations(ctx, db)
}

func (repo *SqlReservationRepository) FindById(ctx context.Context, id uint) (*models.Reservation, error) {
	db, err := getSqlxDatabase(ctx, repo)
	if err != nil {
		return nil, err
	}
	return sqlstore.FindReservationById(ctx, db, id)
}

func (repo *SqlReservationRepository) FindByPlacementId(ctx context.Context, placementId uint) (*models.Reservation, error) {
	db, err := getSqlxDatabase(ctx, repo)
	if err != nil {
		return nil, err
	}
	return sqlstore.FindReservationByPlacementId(ctx, db, placementId)
}

func (repo *SqlReservationRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return inSqlTransaction(ctx, repo, fn)
}
