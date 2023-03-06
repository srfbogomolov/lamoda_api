package unit_test

import (
	"errors"
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
			"Warehouse must be valid",
			models.Warehouse{
				ID:         1,
				Name:       "not empty name",
				IsAvalible: true,
			},
			nil,
		},
		{
			"Warehouse must be not valid",
			models.Warehouse{
				ID:         1,
				Name:       "",
				IsAvalible: true,
			},
			errors.New("warehouse name cannot be empty"),
		},
	}

	for _, tc := range cases {
		err := tc.warehouse.Validate()
		assert.Equal(t, err, tc.expected, tc.desc)
	}
}
