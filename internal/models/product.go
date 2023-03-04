package models

import (
	"context"

	"github.com/google/uuid"
)

type Product struct {
	ID   int       `db:"id"`
	Name string    `db:"name"`
	Size int       `db:"size"`
	Code uuid.UUID `db:"code"`
	QTY  int       `db:"qty"`
}

type ProductRepository interface {
	Save(ctx context.Context, p *Product) error
	FindAll(ctx context.Context) ([]*Product, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}
