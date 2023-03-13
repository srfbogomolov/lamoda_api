package sqlstore

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

// Save available product quantity
func SavePlacement(ctx context.Context, db SqlxDatabase, wDto WarehouseDto, pDto ProductDto) (uint, error) {
	query, _, err := plcPreparedDs.Update().
		Set(goqu.Record{PlacementQTYColumn: "$1"}).
		Where(goqu.And(plcWarehouseIdCol.Eq("$2"), plcProductIdCol.Eq("3"))).
		Returning(plcIdCol).
		ToSQL()
	if err != nil {
		return 0, err
	}
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return 0, err
	}
	var lastId uint
	stmt.GetContext(ctx, &lastId, pDto.AvailableQTY, wDto.Id, pDto.Id)
	return lastId, nil
}
