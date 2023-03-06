package unit_test

import (
	"fmt"
	"testing"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidatePlacement(t *testing.T) {
	cases := []struct {
		desc      string
		placement models.Placement
		expected  error
	}{
		{
			"Placement must be validated",
			models.Placement{
				ProductID:   1,
				WarehouseID: 1,
				QTY:         1,
			},
			nil,
		},
		{
			"Product placement id cannot be less than zero",
			models.Placement{
				ProductID:   0,
				WarehouseID: 1,
				QTY:         1,
			},
			fmt.Errorf("product placement id %w", models.ErrLessOrEqualZero),
		},
		{
			"Warehouse placement id cannot be less than zero",
			models.Placement{
				ProductID:   1,
				WarehouseID: 0,
				QTY:         1,
			},
			fmt.Errorf("warehouse placement id %w", models.ErrLessOrEqualZero),
		},
		{
			"Placement quantity cannot be less than zero",
			models.Placement{
				ProductID:   1,
				WarehouseID: 1,
				QTY:         0,
			},
			fmt.Errorf("placement quantity %w", models.ErrLessOrEqualZero),
		},
	}

	for _, tc := range cases {
		err := tc.placement.Validate()
		assert.Equal(t, tc.expected, err, tc.desc)
	}
}
