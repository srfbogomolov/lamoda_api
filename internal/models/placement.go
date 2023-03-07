package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Placement struct {
	Id          int    `db:"id" json:"id"`
	WarehouseId int    `db:"warehouse_id" json:"warehouse_id"`
	ProductCode string `db:"product_code" json:"product_code"`
	QTY         int    `db:"qty" json:"qty"`
}

func (placement *Placement) Validate() error {
	if _, err := uuid.Parse(placement.ProductCode); err != nil {
		return fmt.Errorf("product placement code %w", ErrIncorrectUUID)
	} else if placement.WarehouseId <= 0 {
		return fmt.Errorf("placement warehouse id %w", ErrLessOrEqualZero)
	} else if placement.QTY <= 0 {
		return fmt.Errorf("placement quantity %w", ErrLessOrEqualZero)
	}
	return nil
}
