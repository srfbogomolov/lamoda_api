package unit_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
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
				ProductCode: uuid.NewString(),
				WarehouseId: 1,
				QTY:         1,
			},
			nil,
		},
		{
			"Product placement code cannot be incorrect",
			models.Placement{
				ProductCode: "111",
				WarehouseId: 1,
				QTY:         1,
			},
			fmt.Errorf("product placement code %w", models.ErrIncorrectUUID),
		},
		{
			"Warehouse placement id cannot be less or equal zero",
			models.Placement{
				ProductCode: uuid.NewString(),
				WarehouseId: 0,
				QTY:         1,
			},
			fmt.Errorf("placement warehouse id %w", models.ErrLessOrEqualZero),
		},
		{
			"Placement quantity cannot be less or equal zero",
			models.Placement{
				ProductCode: uuid.NewString(),
				WarehouseId: 1,
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
