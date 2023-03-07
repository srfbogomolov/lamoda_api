package models

import (
	"fmt"
)

type Warehouse struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	IsAvailable bool   `db:"is_available" json:"is_available"`
}

func (w *Warehouse) Validate() error {
	if w.Name == "" {
		return fmt.Errorf("warehouse name %w", ErrEmpty)
	}
	return nil
}
