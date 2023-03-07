package sqlstore

import (
	"context"
	// dbSql "database/sql"
	// "fmt"

	"github.com/srfbogomolov/warehouse_api/internal/models"
)

const PlacementTable = "placements"

// TODO ПРОТЕСТИРОВАТЬ
// TODO МОЖЕТ ВСЕ-ТАКИ РАЗБИТЬ НА НЕСКОЛЬКО ЗАПРОСОВ?
// func SavePlacementOOOOOOOOOLLLLDDDDD(ctx context.Context, db SqlxDatabase, p *models.Placement) (err error) { // todo убрать еер
// 	warehouse := new(models.Warehouse)
// 	sql := `SELECT * FROM ` + WarehouseTable + ` WHERE id=$1`
// 	if err = db.GetContext(ctx, warehouse, sql, p.WarehouseId); err != nil {
// 		return err
// 	}
// 	if !warehouse.IsAvailable {
// 		return fmt.Errorf("[postgres.AddProductInWarehouse]:warehouse=%v not avalible", warehouse.Id) // TODO ВЕРНУТЬ НОРМ ОШИБКУ, что склад не доступен
// 	}

// 	var productQTY int
// 	sql = `SELECT SUM(qty) FROM ` + PlacementTable + ` WHERE product_code=$1 GROUP BY product_code`
// 	if err = db.GetContext(ctx, &productQTY, sql, p.Id); err != nil {
// 		if err != dbSql.ErrNoRows {
// 			return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err) // TODO ВЕРНУТЬ НОРМ ОШИБКУ, не вернулось не одной строки
// 		}
// 	}

// 	product := new(models.Product)
// 	sql = `SELECT * FROM ` + ProductTable + ` WHERE id=$1`
// 	if err = db.GetContext(ctx, product, sql, p.ProductCode); err != nil {
// 		return err
// 	}
// 	if product.QTY < p.QTY+productQTY {
// 		return fmt.Errorf("[postgres.AddProductInWarehouse]:product.QTY=%v less then data.QTY=%v", product.QTY, p.QTY+productQTY) // TODO ВЕРНУТЬ НОРМ ОШИБКУ, что на складе нет столько товаров
// 	}

// 	sql = `INSERT INTO ` + PlacementTable + ` (product_code, warehouse_id, qty) VALUES($1, $2, $3)
// 			ON CONFLICT ON CONSTRAINT uk_replacement_product_id_warehouse_id DO UPDATE SET "qty"=EXCLUDED.qty+` + PlacementTable + `.qty`
// 	var lastId int
// 	stmt, err := db.PreparexContext(ctx, sql)
// 	if err != nil {
// 		return err
// 	}
// 	stmt.GetContext(ctx, &lastId, p.ProductCode, p.WarehouseId, p.QTY)
// 	p.Id = lastId

// 	return err
// }

func SavePlacement(ctx context.Context, db SqlxDatabase, placement *models.Placement) (int, error) {
	query := `INSERT INTO ` + PlacementTable + ` (warehouse_id, product_code, qty) VALUES ($1, $2, $3)
				ON CONFLICT ON CONSTRAINT uk_replacement_warehouse_id_product_code
					DO UPDATE SET "qty"=EXCLUDED.qty+` + PlacementTable + `.qty RETURNING id`
	var lastId int
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return 0, err
	}
	stmt.GetContext(ctx, &lastId, placement.ProductCode, placement.WarehouseId, placement.QTY)
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
