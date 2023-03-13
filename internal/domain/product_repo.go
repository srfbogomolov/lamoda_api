package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository interface {
	Save(ctx context.Context, product Product) error
	GetTotalByItemId(ctx context.Context, id uint) (Product, error)
	GetTotalByItemCode(ctx context.Context, code uuid.UUID) (Product, error)
	GetInStockByItemId(ctx context.Context, id uint, warehouse_id uint) (Product, error)
	GetInStockByItemCode(ctx context.Context, code uuid.UUID, warehouse_id uint) (Product, error)
	GetAllInStockByWarehouseId(ctx context.Context, warehouse_id uint) (Products, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}
