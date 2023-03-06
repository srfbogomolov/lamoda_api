package models

import (
	"context"
	"fmt"
)

type Product struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Size int    `db:"size" json:"size"`
	Code string `db:"code" json:"code"`
	QTY  int    `db:"qty" json:"qty"`
}

type ProductRepository interface {
	Save(ctx context.Context, p *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("product name %w", ErrEmpty)
	} else if p.Size < 0 {
		return fmt.Errorf("product size %w", ErrLessZero)
	} else if p.QTY <= 0 {
		return fmt.Errorf("product quantity %w", ErrLessOrEqualZero)
	}
	return nil
}
