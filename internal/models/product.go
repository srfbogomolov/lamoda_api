package models

import (
	"fmt"
)

type Product struct {
	Code string `db:"code" json:"code"`
	Name string `db:"name" json:"name"`
	Size int    `db:"size" json:"size"`
	QTY  int    `db:"qty" json:"qty"`
}

func (product *Product) Validate() error {
	if product.Name == "" {
		return fmt.Errorf("product name %w", ErrEmpty)
	} else if product.Size < 0 {
		return fmt.Errorf("product size %w", ErrLessZero)
	} else if product.QTY <= 0 {
		return fmt.Errorf("product quantity %w", ErrLessOrEqualZero)
	}
	return nil
}
