package models

import (
	"context"
	"errors"
)

var (
	ErrEmpty           = errors.New("cannot be empty")
	ErrLessZero        = errors.New("cannot be less than zero")
	ErrLessOrEqualZero = errors.New("cannot be less or equal zero")
)

type Model interface {
	Validate() error
}

type WarehouseRepository interface {
	Save(ctx context.Context, w *Warehouse) error
	FindById(ctx context.Context, id int) (*Warehouse, error)
	Find(ctx context.Context) ([]*Warehouse, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}

type ProductRepository interface {
	Save(ctx context.Context, p *Product) error
	FindByCode(ctx context.Context, code string) (*Product, error)
	Find(ctx context.Context) ([]*Product, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}

type PlacementRepository interface {
	Save(ctx context.Context, p *Placement) error
	FindById(ctx context.Context, id int) (*Placement, error)
	Find(ctx context.Context) ([]*Placement, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}
