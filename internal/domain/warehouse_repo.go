package domain

import (
	"context"
	"errors"
)

var ErrWarehouseNotFound = errors.New("warehouse not found")

type WarehouseRepository interface {
	Save(ctx context.Context, warehouse Warehouse) error
	GetById(ctx context.Context, id uint) (Warehouse, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}
