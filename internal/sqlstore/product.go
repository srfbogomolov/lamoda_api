package sqlstore

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
)

const ProductTable = "products"

func SaveProduct(ctx context.Context, db SqlxDatabase, product *models.Product) (string, error) {
	query := `INSERT INTO ` + ProductTable + ` (name, size, qty) VALUES ($1, $2, $3)
				ON CONFLICT DO NOTHING RETURNING code`
	var lastCode string
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return "", err
	}
	stmt.GetContext(ctx, &lastCode, product.Name, product.Size, product.QTY)
	product.Code = lastCode
	return lastCode, err
}

func FindProductByCode(ctx context.Context, db SqlxDatabase, code string) (*models.Product, error) {
	product := new(models.Product)
	query := `SELECT * FROM ` + ProductTable + ` WHERE code=$1`
	err := db.GetContext(ctx, product, query, code)
	return product, err
}

func FindProducts(ctx context.Context, db SqlxDatabase) ([]*models.Product, error) {
	var products []*models.Product
	query := `SELECT * FROM ` + ProductTable
	err := db.SelectContext(ctx, &products, query)
	return products, err
}
