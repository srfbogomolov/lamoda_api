package sqlstore

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
)

const PlacementTable = "placements"

func SavePlacement(ctx context.Context, db SqlxDatabase, placement *models.Placement) (int, error) {
	query := `INSERT INTO ` + PlacementTable + ` (warehouse_id, product_code, qty) VALUES ($1, $2, $3)
				ON CONFLICT ON CONSTRAINT uk_placement_warehouse_id_product_code
					DO UPDATE SET "qty"=EXCLUDED.qty+` + PlacementTable + `.qty RETURNING id`
	var lastId int
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return 0, err
	}
	stmt.GetContext(ctx, &lastId, placement.WarehouseId, placement.ProductCode, placement.QTY)
	placement.Id = lastId
	return lastId, err
}

func FindPlacementById(ctx context.Context, db SqlxDatabase, id int) (*models.Placement, error) {
	placement := new(models.Placement)
	query := `SELECT * FROM ` + PlacementTable + ` WHERE id=$1`
	err := db.GetContext(ctx, placement, query, id)
	return placement, err
}

func FindPlacements(ctx context.Context, db SqlxDatabase) ([]*models.Placement, error) {
	var placements []*models.Placement
	query := `SELECT * FROM ` + PlacementTable
	err := db.SelectContext(ctx, &placements, query)
	return placements, err
}

func FindPlacementQTYSumByProductCode(ctx context.Context, db SqlxDatabase, productCode string) (int, error) {
	var qtySum int
	query := `SELECT SUM(qty) FROM ` + PlacementTable + ` WHERE product_code = $1 GROUP BY product_code`
	err := db.GetContext(ctx, &qtySum, query, productCode)
	return qtySum, err
}
