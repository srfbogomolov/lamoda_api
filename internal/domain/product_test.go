package domain_test

import (
	"testing"

	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test         string
		name         string
		size         uint
		availableQTY uint
		reservedQTY  uint
		expectedErr  error
	}
	testCases := []testCase{
		{
			test:         "Create new product with not empty item name",
			name:         "test",
			size:         0,
			availableQTY: 0,
			expectedErr:  nil,
		},
		{
			test:         "Create new product with empty item name",
			name:         "",
			size:         0,
			availableQTY: 0,
			expectedErr:  domain.ErrEmptyItemName,
		},
	}

	for _, tc := range testCases {
		_, err := domain.NewProduct(tc.name, tc.size, tc.availableQTY, tc.reservedQTY)
		assert.Equal(t, tc.expectedErr, err, tc.test)
	}
}

func TestProduct_EqualTo(t *testing.T) {
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)
	product, _ := domain.NewProduct("test", 0, 0, 0)
	product.SetItem(item)
	anotherProduct, _ := domain.NewProduct("another test", 0, 0, 0)
	anotherProduct.SetItem(anotherItem)

	type testCase struct {
		test     string
		compared domain.Product
		expected bool
	}
	testCases := []testCase{
		{
			test:     "Product with the same item",
			compared: product,
			expected: true,
		},
		{
			test:     "Product with another item",
			compared: anotherProduct,
			expected: false,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, product.EqualTo(tc.compared), tc.test)
	}
}

func TestProduct_IsOutOfStock(t *testing.T) {
	type testCase struct {
		test         string
		availableQTY uint
		expected     bool
	}
	testCases := []testCase{
		{
			test:         "The available quantity of the product is greater than zero",
			availableQTY: 1,
			expected:     false,
		},
		{
			test:         "The available quantity of the product is zero",
			availableQTY: 0,
			expected:     true,
		},
	}

	for _, tc := range testCases {
		product, _ := domain.NewProduct("test", 0, tc.availableQTY, 0)
		assert.Equal(t, tc.expected, product.IsOutOfStock(), tc.test)
	}
}

func TestProduct_IncreaseAvailableQTY(t *testing.T) {
	var startAvailableQTY uint = 100
	product, _ := domain.NewProduct("test", 0, startAvailableQTY, 0)
	anotherProduct, _ := domain.NewProduct("test", 0, 1, 0)

	type testCase struct {
		test                 string
		summand              domain.Product
		expectedAvailableQTY uint
		expectedErr          error
	}
	testCases := []testCase{
		{
			test:                 "Add a product with the same item",
			summand:              product,
			expectedAvailableQTY: 200,
			expectedErr:          nil,
		},
		{
			test:                 "Add a product with another item",
			summand:              anotherProduct,
			expectedAvailableQTY: 100,
			expectedErr:          domain.ErrProductType,
		},
	}

	for _, tc := range testCases {
		err := product.IncreaseAvailableQTY(tc.summand)
		assert.Equal(t, tc.expectedAvailableQTY, product.GetAvailableQTY(), tc.test)
		assert.Equal(t, tc.expectedErr, err, tc.test)
		product.SetAvailableQTY(startAvailableQTY)
	}
}

func TestProduct_DecreaseAvailableQTY(t *testing.T) {
	var startAvailableQTY uint = 100
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)

	product, _ := domain.NewProduct("test", 0, startAvailableQTY, 0)
	product.SetItem(item)

	productSameItemEnoughQTY, _ := domain.NewProduct(
		"same item enough available qty", 0, 50, 0)
	productSameItemEnoughQTY.SetItem(item)

	productSameItemTooBigQTY, _ := domain.NewProduct(
		"same item too big available qty", 0, 200, 0)
	productSameItemTooBigQTY.SetItem(item)

	productAnotherItemEnoughQTY, _ := domain.NewProduct(
		"another item enough available qty", 0, 50, 0)
	productAnotherItemEnoughQTY.SetItem(anotherItem)

	type testCase struct {
		test                 string
		subtracted           domain.Product
		expectedAvailableQTY uint
		expectedErr          error
	}
	testCases := []testCase{
		{
			test:                 "Same item, enough available quantity to decrease",
			subtracted:           productSameItemEnoughQTY,
			expectedAvailableQTY: 50,
			expectedErr:          nil,
		},
		{
			test:                 "Same item, not enough available quantity to decrease",
			subtracted:           productSameItemTooBigQTY,
			expectedAvailableQTY: startAvailableQTY,
			expectedErr:          domain.ErrNotEnoughAvailableProducts,
		},
		{
			test:                 "Not the same item, enough available quantity to decrease",
			subtracted:           productAnotherItemEnoughQTY,
			expectedAvailableQTY: startAvailableQTY,
			expectedErr:          domain.ErrProductType,
		},
	}

	for _, tc := range testCases {
		err := product.DecreaseAvailableQTY(tc.subtracted)
		assert.Equal(t, tc.expectedAvailableQTY, product.GetAvailableQTY(), tc.test)
		assert.Equal(t, tc.expectedErr, err, tc.test)
		product.SetAvailableQTY(startAvailableQTY)
	}
}

// func TestProduct_IncreaseReservedQTY(t *testing.T) {
// 	var startAvailableQTY uint = 100
// 	var startReservedQTY uint = 0
// 	item, _ := domain.NewItem("test", 0)
// 	product, _ := domain.NewProduct("test", 0, startAvailableQTY, startReservedQTY)
// 	product.SetItem(item)
// 	sameProduct, _ := domain.NewProduct("same test", 0, 0, 50)
// 	sameProduct.SetItem(item)
// 	sameProductTooBigQTY, _ := domain.NewProduct("same test", 0, 0, 200)
// 	sameProductTooBigQTY.SetItem(item)
// 	anotherProduct, _ := domain.NewProduct("test", 0, 0, 1)

// 	type testCase struct {
// 		test                 string
// 		summand              domain.Product
// 		expectedAvailableQTY uint
// 		expectedReservedQTY  uint
// 		expectedErr          error
// 	}
// 	testCases := []testCase{
// 		{
// 			test:                 "Add a reserved product with the same item",
// 			summand:              sameProduct,
// 			expectedAvailableQTY: 50,
// 			expectedReservedQTY:  50,
// 			expectedErr:          nil,
// 		},
// 		{
// 			test:                 "Add a reserved product with another item",
// 			summand:              anotherProduct,
// 			expectedAvailableQTY: startAvailableQTY,
// 			expectedReservedQTY:  startReservedQTY,
// 			expectedErr:          domain.ErrProductType,
// 		},
// 		{
// 			test:                 "Add too many reserved products",
// 			summand:              sameProductTooBigQTY,
// 			expectedAvailableQTY: startAvailableQTY,
// 			expectedReservedQTY:  startReservedQTY,
// 			expectedErr:          domain.ErrNotEnoughAvailableProducts,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		err := product.IncreaseReservedQTY(tc.summand)
// 		assert.Equal(t, tc.expectedAvailableQTY, product.GetAvailableQTY(), tc.test)
// 		assert.Equal(t, tc.expectedReservedQTY, product.GetReservedQTY(), tc.test)
// 		assert.Equal(t, tc.expectedErr, err, tc.test)
// 		product.SetAvailableQTY(startAvailableQTY)
// 		product.SetReservedQTY(startReservedQTY)
// 	}
// }

func TestProduct_DecreaseReservedQTY(t *testing.T) {
	var startAvailableQTY uint = 0
	var startReservedQTY uint = 100
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)

	product, _ := domain.NewProduct("test", 0, startAvailableQTY, startReservedQTY)
	product.SetItem(item)

	productSameItemEnoughQTY, _ := domain.NewProduct(
		"same item enough reserved qty", 0, 0, 50)
	productSameItemEnoughQTY.SetItem(item)

	productSameItemTooBigQTY, _ := domain.NewProduct(
		"same item too big reserved qty", 0, 0, 200)
	productSameItemTooBigQTY.SetItem(item)

	productAnotherItemEnoughQTY, _ := domain.NewProduct(
		"another item enough reserved qty", 0, 0, 50)
	productAnotherItemEnoughQTY.SetItem(anotherItem)

	type testCase struct {
		test                 string
		subtracted           domain.Product
		expectedAvailableQTY uint
		expectedReservedQTY  uint
		expectedErr          error
	}
	testCases := []testCase{
		{
			test:                 "Same item, enough reserved quantity to decrease",
			subtracted:           productSameItemEnoughQTY,
			expectedAvailableQTY: 50,
			expectedReservedQTY:  50,
			expectedErr:          nil,
		},
		{
			test:                 "Same item, not enough reserved quantity to decrease",
			subtracted:           productSameItemTooBigQTY,
			expectedAvailableQTY: startAvailableQTY,
			expectedReservedQTY:  startReservedQTY,
			expectedErr:          domain.ErrNotEnoughReservedProducts,
		},
		{
			test:                 "Not the same item, enough reserved quantity to decrease",
			subtracted:           productAnotherItemEnoughQTY,
			expectedAvailableQTY: startAvailableQTY,
			expectedReservedQTY:  startReservedQTY,
			expectedErr:          domain.ErrProductType,
		},
	}

	for _, tc := range testCases {
		err := product.DecreaseReservedQTY(tc.subtracted)
		assert.Equal(t, tc.expectedAvailableQTY, product.GetAvailableQTY(), tc.test)
		assert.Equal(t, tc.expectedReservedQTY, product.GetReservedQTY(), tc.test)
		assert.Equal(t, tc.expectedErr, err, tc.test)
		product.SetAvailableQTY(startAvailableQTY)
		product.SetReservedQTY(startReservedQTY)
	}
}

func TestProducts_IsOutOfStock(t *testing.T) {
	productWithZeroAvailableQTY, _ := domain.NewProduct("test", 0, 0, 0)
	productWithNonZeroAvailableQTY, _ := domain.NewProduct("test", 0, 1, 0)

	type testCase struct {
		test     string
		products domain.Products
		expected bool
	}
	testCases := []testCase{
		{
			test:     "No products at all",
			products: domain.Products{},
			expected: false,
		},
		{
			test: "There is a product, but its available quantity is zero",
			products: domain.Products{
				productWithZeroAvailableQTY,
			},
			expected: false,
		},
		{
			test: "There is a product with a non-zero available quantity",
			products: domain.Products{
				productWithNonZeroAvailableQTY,
			},
			expected: true,
		},
		{
			test: "There are a products with different available quantities",
			products: domain.Products{
				productWithZeroAvailableQTY,
				productWithNonZeroAvailableQTY,
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.products.IsOutOfStock(), tc.test)
	}
}

func TestProducts_IncreaseAvailableQTY(t *testing.T) {
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)
	product, _ := domain.NewProduct("test", 0, 100, 0)
	product.SetItem(item)
	products := domain.Products{product}
	productIncreased, _ := domain.NewProduct("test", 0, 200, 0)
	productIncreased.SetItem(item)
	productAnotherItemZeroQTY, _ := domain.NewProduct("test", 0, 0, 0)
	productAnotherItemNonZeroQTY, _ := domain.NewProduct("test", 0, 100, 0)
	productAnotherItemNonZeroQTY.SetItem(anotherItem)
	productsIncreased := domain.Products{
		productIncreased,
		productAnotherItemNonZeroQTY,
	}

	type testCase struct {
		test             string
		summand          domain.Products
		expectedProducts domain.Products
	}
	testCases := []testCase{
		{
			test:             "Add no products",
			summand:          domain.Products{},
			expectedProducts: products,
		},
		{
			test: "Add products with zero available quantities",
			summand: domain.Products{
				productAnotherItemZeroQTY,
			},
			expectedProducts: products,
		},
		{
			test: "Add different products with zero and non-zero available quantities",
			summand: domain.Products{
				product,
				productAnotherItemZeroQTY,
				productAnotherItemNonZeroQTY,
			},
			expectedProducts: productsIncreased,
		},
	}

	for _, tc := range testCases {
		resultProducts := products.IncreaseAvailableQTY(tc.summand)
		assert.Equal(t, tc.expectedProducts, resultProducts, tc.test)
	}
}

func TestProducts_DecreaseAvailableQTY(t *testing.T) {
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)
	product, _ := domain.NewProduct("test", 0, 100, 0)
	product.SetItem(item)
	productDecreased, _ := domain.NewProduct("test", 0, 0, 0)
	productDecreased.SetItem(item)
	products := domain.Products{product}
	productAnotherItemZeroQTY, _ := domain.NewProduct("test", 0, 0, 0)
	productAnotherItemNonZeroQTY, _ := domain.NewProduct("test", 0, 1, 0)
	productAnotherItemNonZeroQTY.SetItem(anotherItem)

	type testCase struct {
		test             string
		subtracted       domain.Products
		expectedProducts domain.Products
		expectedErr      error
	}
	testCases := []testCase{
		{
			test: "Remove products with zero available quantities",
			subtracted: domain.Products{
				productAnotherItemZeroQTY,
			},
			expectedProducts: domain.Products{
				product,
			},
			expectedErr: nil,
		},
		{
			test: "Remove different products with zero and non-zero available quantities",
			subtracted: domain.Products{
				product,
				productAnotherItemZeroQTY,
				productAnotherItemNonZeroQTY,
			},
			expectedProducts: products,
			expectedErr:      domain.ErrProductType,
		},
		{
			test: "Remove products with same item and non-zero available quantities",
			subtracted: domain.Products{
				product,
				productAnotherItemZeroQTY,
			},
			expectedProducts: domain.Products{
				productDecreased,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		resultProducts, err := products.DecreaseAvailableQTY(tc.subtracted)
		assert.Equal(t, tc.expectedProducts, resultProducts, tc.test)
		assert.Equal(t, tc.expectedErr, err, tc.test)
	}
}

func TestProducts_IncreaseReservedQTY(t *testing.T) {
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)
	product, _ := domain.NewProduct("test", 0, 100, 0)
	product.SetItem(item)
	products := domain.Products{product}
	productReservedIncreased, _ := domain.NewProduct("test", 0, 0, 100)
	productReservedIncreased.SetItem(item)
	productSameItemZeroQTY, _ := domain.NewProduct("test", 0, 0, 0)
	productSameItemZeroQTY.SetItem(item)
	productAnotherItemNonZeroQTY, _ := domain.NewProduct("test", 0, 100, 0)
	productAnotherItemNonZeroQTY.SetItem(anotherItem)
	productsReservedIncreased := domain.Products{
		productReservedIncreased,
	}

	type testCase struct {
		test             string
		summand          domain.Products
		expectedProducts domain.Products
		expectedErr      error
	}
	testCases := []testCase{
		{
			test:             "No products",
			summand:          domain.Products{},
			expectedProducts: products,
			expectedErr:      nil,
		},
		{
			test: "Too many product types",
			summand: domain.Products{
				domain.Product{},
				domain.Product{},
			},
			expectedProducts: nil,
			expectedErr:      domain.ErrNoProductItemInStock,
		},
		{
			test: "Products with another items",
			summand: domain.Products{
				productAnotherItemNonZeroQTY,
			},
			expectedProducts: nil,
			expectedErr:      domain.ErrNoProductItemInStock,
		},
		{
			test: "Add products with zero reserved quantities",
			summand: domain.Products{
				productSameItemZeroQTY,
			},
			expectedProducts: products,
			expectedErr:      nil,
		},
		{
			test: "Add different products with zero and non-zero reserved quantities",
			summand: domain.Products{
				product,
				productAnotherItemNonZeroQTY,
			},
			expectedProducts: nil,
			expectedErr:      domain.ErrNoProductItemInStock,
		},
		{
			test: "Add same products with non-zero reserved quantities",
			summand: domain.Products{
				productReservedIncreased,
			},
			expectedProducts: productsReservedIncreased,
			expectedErr:      nil,
		},
	}

	for _, tc := range testCases {
		resultProducts, err := products.IncreaseReservedQTY(tc.summand)
		assert.Equal(t, tc.expectedProducts, resultProducts, tc.test)
		assert.Equal(t, tc.expectedErr, err, tc.test)
	}
}

func TestProducts_DecreaseReservedQTY(t *testing.T) {
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)
	product, _ := domain.NewProduct("test", 0, 0, 100)
	product.SetItem(item)
	productReservedDecreased, _ := domain.NewProduct("test", 0, 100, 0)
	productReservedDecreased.SetItem(item)
	products := domain.Products{product}
	productSameItemZeroQTY, _ := domain.NewProduct("test", 0, 0, 0)
	productSameItemZeroQTY.SetItem(item)
	productAnotherItemNonZeroQTY, _ := domain.NewProduct("test", 0, 1, 0)
	productAnotherItemNonZeroQTY.SetItem(anotherItem)
	productsReservedDecreased := domain.Products{
		product,
	}

	type testCase struct {
		test             string
		subtracted       domain.Products
		expectedProducts domain.Products
		expectedErr      error
	}
	testCases := []testCase{
		{
			test:             "Remove no products",
			subtracted:       domain.Products{},
			expectedProducts: products,
			expectedErr:      nil,
		},
		{
			test: "Remove products with zero reserved quantities",
			subtracted: domain.Products{
				productSameItemZeroQTY,
			},
			expectedProducts: domain.Products{
				product,
			},
			expectedErr: nil,
		},
		{
			test: "Remove too many product types",
			subtracted: domain.Products{
				product,
				productSameItemZeroQTY,
				productAnotherItemNonZeroQTY,
			},
			expectedProducts: nil,
			expectedErr:      domain.ErrNoProductItemInStock,
		},
		{
			test: "Remove another product",
			subtracted: domain.Products{
				productAnotherItemNonZeroQTY,
			},
			expectedProducts: nil,
			expectedErr:      domain.ErrNoProductItemInStock,
		},
		{
			test: "Remove same products with non-zero reserved quantities",
			subtracted: domain.Products{
				productReservedDecreased,
			},
			expectedProducts: productsReservedDecreased,
			expectedErr:      nil,
		},
	}

	for _, tc := range testCases {
		resultProducts, err := products.DecreaseReservedQTY(tc.subtracted)
		assert.Equal(t, tc.expectedProducts, resultProducts, tc.test)
		assert.Equal(t, tc.expectedErr, err, tc.test)
	}
}

func TestProducts_Contain(t *testing.T) {
	item, _ := domain.NewItem("test", 0)
	anotherItem, _ := domain.NewItem("another test", 0)
	product, _ := domain.NewProduct("test", 0, 0, 0)
	product.SetItem(item)
	anotherProduct, _ := domain.NewProduct("another test", 0, 0, 0)
	anotherProduct.SetItem(anotherItem)
	products := domain.Products{product}

	type testCase struct {
		test     string
		product  domain.Product
		expected bool
	}
	testCases := []testCase{
		{
			test:     "Products contain product",
			product:  product,
			expected: true,
		},
		{
			test:     "Products do not contain product",
			product:  anotherProduct,
			expected: false,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, products.Contain(tc.product), tc.test)
	}
}
