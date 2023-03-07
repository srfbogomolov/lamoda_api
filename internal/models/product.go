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
