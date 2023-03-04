package models

import "context"

type Warehouse struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	IsAvalible bool   `db:"is_available"`
}

type WarehouseRepository interface {
	Save(ctx context.Context, w *Warehouse) error
	FindAll(ctx context.Context) ([]*Warehouse, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}
