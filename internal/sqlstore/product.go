package sqlstore

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
)

const ProductTable = "products"

func SaveProduct(ctx context.Context, db SqlxDatabase, p *models.Product) error {
	sql := `INSERT INTO ` + ProductTable + `(name, size, qty) VALUES($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id`
	var lastId int
	stmt, err := db.PreparexContext(ctx, sql)
	if err != nil {
		return err
	}
	stmt.GetContext(ctx, &lastId, p.Name, p.Size, p.QTY)
	p.ID = lastId
	return err
}

func GetAllProduct(ctx context.Context, db SqlxDatabase) ([]*models.Product, error) {
	var comments []*models.Product
	sql := `SELECT * FROM ` + ProductTable
	err := db.SelectContext(ctx, &comments, sql)
	return comments, err
}
