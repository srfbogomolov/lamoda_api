package sqlstore

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

// Save reserved product quantity
func SaveReservation(ctx context.Context, db SqlxDatabase, placement_id uint, pDto ProductDto) error {
	query, _, err := rsvPreparedDs.Insert().
		Rows(goqu.Record{ReservationPlacementIdColumn: "$1", ReservationQTYColumn: "$2"}).
		OnConflict(goqu.DoUpdate(ReservationPlacementIdColumn, goqu.Record{ReservationQTYColumn: goqu.L("EXCLUDED.qty")})).ToSQL()

	if err != nil {
		return err
	}
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}
	var lastId uint
	stmt.GetContext(ctx, &lastId, placement_id, pDto.ReservedQTY)
	return nil
}
