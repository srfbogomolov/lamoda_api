package domain_test

import (
	"testing"

	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestWarehouse_NewWarehouse(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		isAvailable bool
		products    domain.Products
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "Create new warehouse with not empty name",
			name:        "test",
			isAvailable: true,
			products:    domain.Products{},
			expectedErr: nil,
		},
		{
			test:        "Create new warehouse with empty name",
			name:        "",
			isAvailable: true,
			products:    domain.Products{},
			expectedErr: domain.ErrEmptyWarehouseName,
		},
	}

	for _, tc := range testCases {
		_, err := domain.NewWarehouse(tc.name, tc.isAvailable, tc.products)
		assert.Equal(t, tc.expectedErr, err, tc.test)
	}
}

func TestWarehouse_IsAvailable(t *testing.T) {
	products := domain.Products{}
	availableWarehouse, _ := domain.NewWarehouse("test", true, products)
	notAvailableWarehouse, _ := domain.NewWarehouse("test", false, products)
	type testCase struct {
		test      string
		warehouse domain.Warehouse
		expected  bool
	}
	testCases := []testCase{
		{
			test:      "Warehouse is available",
			warehouse: availableWarehouse,
			expected:  true,
		},
		{
			test:      "Warehouse is not available",
			warehouse: notAvailableWarehouse,
			expected:  false,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.warehouse.IsAvailable(), tc.test)
	}
}

func TestWarehouse_ReserveProducts(t *testing.T) {
	item, _ := domain.NewItem("test", 0)
	product, _ := domain.NewProduct("test", 0, 100, 0)
	product.SetItem(item)
	sameProduct, _ := domain.NewProduct("test", 0, 0, 50)
	sameProduct.SetItem(item)
	anotherProduct, _ := domain.NewProduct("test", 0, 0, 50)
	products := domain.Products{product}
	warehouse, _ := domain.NewWarehouse("test", true, nil)
	warehouse.SetProducts(products)
	notAvailableWarehouse, _ := domain.NewWarehouse("test", false, nil)

	type testCase struct {
		test        string
		products    domain.Products
		warehouse   domain.Warehouse
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "Warehouse is not available",
			products:    domain.Products{},
			warehouse:   notAvailableWarehouse,
			expectedErr: domain.ErrWarehouseIsNotAvailable,
		},
		{
			test: "Reserve a product that is in stock",
			products: domain.Products{
				sameProduct,
			},
			warehouse:   warehouse,
			expectedErr: nil,
		},
		{
			test: "Reserve a product that is not in stock",
			products: domain.Products{
				anotherProduct,
			},
			warehouse:   warehouse,
			expectedErr: domain.ErrNoProductItemInStock,
		},
	}

	for _, tc := range testCases {
		err := tc.warehouse.Reserve(tc.products)
		assert.Equal(t, tc.expectedErr, err, tc.test)
	}
}

func TestWarehouse_RemoveProducts(t *testing.T) {
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)
	product, _ := domain.NewProduct("test", 0, 0, 100)
	product.SetItem(item)
	// productIncreased, _ := domain.NewProduct("test", 0, 2)
	// productIncreased.SetItem(item)
	productWithZeroQTY, _ := domain.NewProduct("test", 0, 0, 0)
	productAnotherItemNonZeroQTY, _ := domain.NewProduct("test", 0, 1, 0)
	productAnotherItemNonZeroQTY.SetItem(anotherItem)
	products := domain.Products{product}
	warehouse, _ := domain.NewWarehouse("test", true, nil)
	warehouse.SetProducts(products)
	notAvailableWarehouse, _ := domain.NewWarehouse("test", false, nil)

	type testCase struct {
		test        string
		products    domain.Products
		warehouse   domain.Warehouse
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "Warehouse is not available",
			products:    domain.Products{},
			warehouse:   notAvailableWarehouse,
			expectedErr: domain.ErrWarehouseIsNotAvailable,
		},
		{
			test: "",
			products: domain.Products{
				product,
			},
			warehouse:   warehouse,
			expectedErr: nil,
		},
		{
			test: "No products to remove from the warehouse",
			products: domain.Products{
				productWithZeroQTY,
			},
			warehouse:   warehouse,
			expectedErr: domain.ErrNoProductItemInStock,
		},
		{
			test: "Remove different products from the warehouse",
			products: domain.Products{
				product,
				productWithZeroQTY,
				productAnotherItemNonZeroQTY,
			},
			warehouse:   warehouse,
			expectedErr: domain.ErrNoProductItemInStock,
		},
	}

	for _, tc := range testCases {
		err := tc.warehouse.Release(tc.products)
		assert.Equal(t, tc.expectedErr, err, tc.test)
	}
}
