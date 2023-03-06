package unit_test

import (
	"fmt"
	"testing"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateWarehouse(t *testing.T) {
	cases := []struct {
		desc      string
		warehouse models.Warehouse
		expected  error
	}{
		{
			"Warehouse must be validated",
			models.Warehouse{
				ID:         0,
				Name:       "not empty name",
				IsAvalible: false,
			},
			nil,
		},
		{
			"Warehouse name cannot be empty",
			models.Warehouse{
				ID:         0,
				Name:       "",
				IsAvalible: false,
			},
			fmt.Errorf("warehouse name %w", errEmpty),
		},
	}

	for _, tc := range cases {
		err := tc.warehouse.Validate()
		assert.Equal(t, tc.expected, err, tc.desc)
	}
}
