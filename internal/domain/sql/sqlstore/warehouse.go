package sqlstore

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
)

type WarehouseDto struct {
	Id          uint   `db:"id"`
	Name        string `db:"name"`
	IsAvailable bool   `db:"is_available"`
}

// Create warehouse from dto
func (d WarehouseDto) ToAggregate() domain.Warehouse {
	w := domain.Warehouse{}
	w.SetId(d.Id)
	w.SetName(d.Name)
	w.SetIsAvailable(d.IsAvailable)

	return w
}

// Create dto from warehouse
func FromWarehouse(w domain.Warehouse) WarehouseDto {
	return WarehouseDto{
		Id:          w.GetId(),
		Name:        w.GetName(),
		IsAvailable: w.GetIsAvailable(),
	}
}

// Save warehouse
func SaveWarehouse(ctx context.Context, db SqlxDatabase, dto WarehouseDto) error {
	sql, _, err := wrhsPreparedDs.Update().
		Set(goqu.Record{WarehouseNameColumn: "$1", WarehouseIsAvailableColumn: "$2"}).
		Where(wrhsIdCol.Eq("$3")).
		ToSQL()
	if err != nil {
		return err
	}
	stmt, err := db.PreparexContext(ctx, sql)
	if err != nil {
		return err
	}
	var lastId uint
	stmt.GetContext(ctx, &lastId, dto.IsAvailable, dto.Name, dto.Id)
	return nil
}

// Get warehouse by id
func GetWarehouseById(ctx context.Context, db SqlxDatabase, id uint) (domain.Warehouse, error) {
	query, _, err := wrhsPreparedDs.Where(wrhsIdCol.Eq("$1")).ToSQL()
	if err != nil {
		return domain.Warehouse{}, err
	}
	dto := new(WarehouseDto)
	err = db.GetContext(ctx, dto, query, id)
	if err != nil {
		return domain.Warehouse{}, domain.ErrWarehouseNotFound
	}
	return dto.ToAggregate(), nil
}
