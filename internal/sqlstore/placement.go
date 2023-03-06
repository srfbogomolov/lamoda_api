package sqlstore

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
)

const PlacementTable = "placements"

func SavePlacement(ctx context.Context, db SqlxDatabase, p *models.Placement) error {
	sql := `INSERT INTO ` + PlacementTable + `(product_id, warehouse_id, qty) VALUES($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id`
	var lastId int
	stmt, err := db.PreparexContext(ctx, sql)
	if err != nil {
		return err
	}
	stmt.GetContext(ctx, &lastId, p.ProductID, p.WarehouseID, p.QTY)
	p.ID = lastId
	return err
}

func GetPlacementByID(ctx context.Context, db SqlxDatabase, id int) (*models.Placement, error) {
	p := new(models.Placement)
	sql := `SELECT * FROM ` + PlacementTable + ` WHERE id=$1`
	err := db.GetContext(ctx, &p, sql, id)
	return p, err
}

func GetAllPlacements(ctx context.Context, db SqlxDatabase) ([]*models.Placement, error) {
	var comments []*models.Placement
	sql := `SELECT * FROM ` + PlacementTable
	err := db.SelectContext(ctx, &comments, sql)
	return comments, err
}
