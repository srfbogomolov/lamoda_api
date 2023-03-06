package unit_test

import (
	"fmt"
	"testing"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateProduct(t *testing.T) {
	cases := []struct {
		desc     string
		product  models.Product
		expected error
	}{
		{
			"Product must be validated",
			models.Product{
				ID:   1,
				Name: "not empty name",
				Size: 0,
				QTY:  1,
			},
			nil,
		},
		{
			"Product name cannot be empty",
			models.Product{
				ID:   1,
				Name: "",
				Size: 0,
				QTY:  1,
			},
			fmt.Errorf("product name %w", models.ErrEmpty),
		},
		{
			"Product size cannot be less than zero",
			models.Product{
				ID:   1,
				Name: "not empty name",
				Size: -1,
				QTY:  1,
			},
			fmt.Errorf("product size %w", models.ErrLessZero),
		},
		{
			"Product quantity cannot be less than zero",
			models.Product{
				ID:   1,
				Name: "not empty name",
				Size: 0,
				QTY:  0,
			},
			fmt.Errorf("product quantity %w", models.ErrLessOrEqualZero),
		},
	}

	for _, tc := range cases {
		err := tc.product.Validate()
		assert.Equal(t, tc.expected, err, tc.desc)
	}
}
