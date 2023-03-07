package sqlstore

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
)

const WarehouseTable = "warehouses"

func SaveWarehouse(ctx context.Context, db SqlxDatabase, warehouse *models.Warehouse) (int, error) {
	query := `INSERT INTO ` + WarehouseTable + `(name, is_available) VALUES ($1, $2)
				ON CONFLICT DO NOTHING RETURNING id`
	var lastId int
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return 0, err
	}
	stmt.GetContext(ctx, &lastId, warehouse.Name, warehouse.IsAvailable)
	warehouse.Id = lastId
	return lastId, err
}

func FindWarehouseById(ctx context.Context, db SqlxDatabase, id int) (*models.Warehouse, error) {
	warehouse := new(models.Warehouse)
	query := `SELECT * FROM ` + WarehouseTable + ` WHERE id=$1`
	err := db.GetContext(ctx, warehouse, query, id)
	return warehouse, err
}

func FindWarehouses(ctx context.Context, db SqlxDatabase) ([]*models.Warehouse, error) {
	var warehouses []*models.Warehouse
	query := `SELECT * FROM ` + WarehouseTable
	err := db.SelectContext(ctx, &warehouses, query)
	return warehouses, err
}
