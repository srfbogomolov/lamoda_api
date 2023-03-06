package unit_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
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
				ID:   0,
				Name: "not empty name",
				Size: 0,
				Code: uuid.NewString(),
				QTY:  0,
			},
			nil,
		},
		{
			"Product name cannot be empty",
			models.Product{
				ID:   0,
				Name: "",
				Size: 0,
				Code: uuid.NewString(),
				QTY:  0,
			},
			fmt.Errorf("product name %w", errEmpty),
		},
		{
			"Product size cannot be less than zero",
			models.Product{
				ID:   0,
				Name: "not empty name",
				Size: -1,
				Code: uuid.NewString(),
				QTY:  0,
			},
			fmt.Errorf("product size %w", errLessThanZero),
		},
		{
			"Product quantity cannot be less than zero",
			models.Product{
				ID:   0,
				Name: "not empty name",
				Size: 0,
				Code: uuid.NewString(),
				QTY:  -1,
			},
			fmt.Errorf("product quantity %w", errLessThanZero),
		},
	}

	for _, tc := range cases {
		err := tc.product.Validate()
		assert.Equal(t, tc.expected, err, tc.desc)
	}
}
