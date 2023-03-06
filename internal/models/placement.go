package models

import (
	"context"
	"fmt"
)

type Placement struct {
	ID          int `db:"id" json:"id"`
	ProductID   int `db:"product_id" json:"product_id"`
	WarehouseID int `db:"warehouse_id" json:"warehouse_id"`
	QTY         int `db:"qty" json:"qty"`
}

type PlacementRepository interface {
	Save(ctx context.Context, p *Placement) error
	GetByID(ctx context.Context, id int) (*Placement, error)
	GetAll(ctx context.Context) ([]*Placement, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}

func (p *Placement) Validate() error {
	if p.ProductID <= 0 {
		return fmt.Errorf("product placement id %w", ErrLessOrEqualZero)
	} else if p.WarehouseID <= 0 {
		return fmt.Errorf("warehouse placement id %w", ErrLessOrEqualZero)
	} else if p.QTY <= 0 {
		return fmt.Errorf("placement quantity %w", ErrLessOrEqualZero)
	}
	return nil
}
