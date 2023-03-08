package sqlstore

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/srfbogomolov/warehouse_api/internal/models"
)

const ReservationTable = "reservations"

var (
	reservationT              = goqu.T(ReservationTable)
	reservationIdCol          = reservationT.Col("id")
	reservationPlacementIdCol = reservationT.Col("placement_id")
	reservationQTYCol         = reservationT.Col("qty")
	reservationPreparedDs     = dialect.From(reservationT).Prepared(true)
)

func SaveReservation(ctx context.Context, db SqlxDatabase, reservation *models.Reservation) (uint, error) {
	sql, _, err := reservationPreparedDs.Insert().
		Cols(goqu.C("placement_id"), goqu.C("qty")).
		Vals(goqu.Vals{reservation.PlacementId, reservation.QTY}).
		Returning("id").
		ToSQL()
	if err != nil {
		return 0, err
	}
	stmt, err := db.PreparexContext(ctx, sql)
	if err != nil {
		return 0, err
	}
	var lastId uint
	stmt.GetContext(ctx, &lastId, reservation.PlacementId, reservation.QTY)
	reservation.Id = lastId
	return lastId, err
}

func FindReservations(ctx context.Context, db SqlxDatabase) ([]*models.Reservation, error) {
	sql, _, err := reservationPreparedDs.ToSQL()
	if err != nil {
		return nil, err
	}
	var reservations []*models.Reservation
	err = db.SelectContext(ctx, &reservations, sql)
	return reservations, err
}

func FindReservationById(ctx context.Context, db SqlxDatabase, id int) (*models.Reservation, error) {
	sql, _, err := reservationPreparedDs.Where(reservationIdCol.Eq("$1")).ToSQL()
	if err != nil {
		return nil, err
	}
	reservation := new(models.Reservation)
	err = db.GetContext(ctx, reservation, sql, id)
	return reservation, err
}

func FindReservationByPlacementId(ctx context.Context, db SqlxDatabase, placementId int) (*models.Reservation, error) {
	sql, _, err := reservationPreparedDs.Where(reservationPlacementIdCol.Eq("$1")).ToSQL()
	if err != nil {
		return nil, err
	}
	reservation := new(models.Reservation)
	err = db.GetContext(ctx, reservation, sql, placementId)
	return reservation, err
}
