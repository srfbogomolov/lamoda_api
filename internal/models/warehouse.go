package models

import (
	"context"
	"errors"
)

var errEmptyWarehouseName = errors.New("warehouse name cannot be empty")

type Warehouse struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	IsAvalible bool   `db:"is_available" json:"is_available"`
}

type WarehouseRepository interface {
	Save(ctx context.Context, w *Warehouse) error
	GetAll(ctx context.Context) ([]*Warehouse, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}

func (w *Warehouse) Validate() error {
	if w.Name == "" {
		return errEmptyWarehouseName
	}
	return nil
}
