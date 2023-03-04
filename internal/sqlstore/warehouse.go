package sqlstore

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
)

const WarehouseTable = "warehouses"

func SaveWarehouse(ctx context.Context, db SqlxDatabase, w *models.Warehouse) error {
	sql := `INSERT INTO ` + WarehouseTable + `(name, is_available) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING id`
	var lastId int
	stmt, err := db.PreparexContext(ctx, sql)
	if err != nil {
		return err
	}
	stmt.GetContext(ctx, &lastId, w.Name, w.IsAvalible)
	w.ID = lastId
	return err
}

func FindAllWarehouse(ctx context.Context, db SqlxDatabase) ([]*models.Warehouse, error) {
	var warehouses []*models.Warehouse
	sql := `SELECT * FROM ` + WarehouseTable
	err := db.SelectContext(ctx, &warehouses, sql)

	return warehouses, err
}
