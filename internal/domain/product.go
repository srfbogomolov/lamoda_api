package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotEnoughAvailableProducts = errors.New("not enough available products")
	ErrNotEnoughReservedProducts  = errors.New("not enough reserved products")
	ErrProductType                = errors.New("product type does not match")
	ErrProductIsOutOfStock        = errors.New("product is out of stock")
	ErrNoProductItemInStock       = errors.New("there are no products of this type in stock")
)

// Quantity of a specific type of item
type Product struct {
	item         *Item
	availableQTY uint
	reservedQTY  uint
}

type Products []Product

// Create new product
func NewProduct(name string, size uint, availableQTY uint, reservedQTY uint) (Product, error) { // reservedQTY
	item, err := NewItem(name, size)
	if err != nil {
		return Product{}, err
	}
	return Product{
		item:         item,
		availableQTY: availableQTY,
		reservedQTY:  reservedQTY,
	}, nil
}

// Get product item id
func (p Product) GetId() uint {
	return p.item.Id
}

// Set product item id
func (p *Product) SetId(id uint) {
	p.item.Id = id
}

// Get product item code
func (p Product) GetCode() uuid.UUID {
	return p.item.Code
}

// Set product item code
func (p *Product) SetCode(code uuid.UUID) {
	p.item.Code = code
}

// Get product item name
func (p Product) GetName() string {
	return p.item.Name
}

// Set product item name
func (p *Product) SetName(name string) {
	p.item.Name = name
}

// Get product item size
func (p Product) GetSize() uint {
	return p.item.Size
}

// Set product item size
func (p *Product) SetSize(size uint) {
	p.item.Size = size
}

// Get product item
func (p Product) GetItem() *Item {
	return p.item
}

// Set product item
func (p *Product) SetItem(item *Item) {
	p.item = item
}

// Get product available quantity
func (p Product) GetAvailableQTY() uint {
	return p.availableQTY
}

// Set product available quantity
func (p *Product) SetAvailableQTY(availableQTY uint) {
	p.availableQTY = availableQTY
}

// Get product reserved quantity
func (p Product) GetReservedQTY() uint {
	return p.reservedQTY
}

// Set product reserved quantity
func (p *Product) SetReservedQTY(reservedQTY uint) {
	p.reservedQTY = reservedQTY
}

// Check product is equal to another product
func (p Product) EqualTo(other Product) bool {
	return p.item.EqualTo(*other.item)
}

// Check product is out of stock or not
func (p Product) IsOutOfStock() bool {
	return p.availableQTY == 0 && p.reservedQTY == 0
}

// Increase available quantity of specific type of item
func (p *Product) IncreaseAvailableQTY(other Product) error {
	if !other.item.EqualTo(*p.item) {
		return ErrProductType
	}
	p.availableQTY = p.availableQTY + other.availableQTY
	return nil
}

// Decrease available quantity of specific type of item
func (p *Product) DecreaseAvailableQTY(other Product) error {
	if !other.item.EqualTo(*p.item) {
		return ErrProductType
	}
	if other.availableQTY > p.availableQTY {
		return ErrNotEnoughAvailableProducts
	}
	p.availableQTY = p.availableQTY - other.availableQTY
	return nil
}

// Increase reserved quantity of specific type of item
func (p *Product) IncreaseReservedQTY(other Product) error {
	if !other.item.EqualTo(*p.item) {
		return ErrProductType
	} else if other.reservedQTY > p.availableQTY {
		return ErrNotEnoughAvailableProducts
	} else if other.IsOutOfStock() {
		return nil
	}
	p.availableQTY = p.availableQTY - other.reservedQTY
	p.reservedQTY = p.reservedQTY + other.reservedQTY
	return nil
}

// Decrease reserved quantity of specific type of item
func (p *Product) DecreaseReservedQTY(other Product) error {
	if !other.item.EqualTo(*p.item) {
		return ErrProductType
	}
	if other.reservedQTY > p.reservedQTY {
		return ErrNotEnoughReservedProducts
	}
	p.availableQTY = p.availableQTY + other.reservedQTY
	p.reservedQTY = p.reservedQTY - other.reservedQTY
	return nil
}

// Check products are in stock or not
func (ps Products) IsOutOfStock() bool {
	for _, p := range ps {
		if !p.IsOutOfStock() {
			return true
		}
	}
	return false
}

// Increase available quantities of products
func (ps Products) IncreaseAvailableQTY(other Products) Products {
	if !other.IsOutOfStock() {
		return ps
	}

	result := Products{}
	for _, psProduct := range ps {
		for _, otherProduct := range other {
			if otherProduct.IsOutOfStock() {
				continue
			} else if err := psProduct.IncreaseAvailableQTY(otherProduct); err != nil {
				result = append(result, otherProduct)
			} else {
				result = append(result, psProduct)
			}
		}
	}

	return result
}

// Decrease available quantities of products
func (ps Products) DecreaseAvailableQTY(other Products) (Products, error) {
	if !other.IsOutOfStock() {
		return ps, nil
	}

	res := Products{}
	for _, psProduct := range ps {
		for _, otherProduct := range other {
			if otherProduct.IsOutOfStock() {
				continue
			}
			if err := psProduct.DecreaseAvailableQTY(otherProduct); err != nil {
				return ps, err
			}
			res = append(res, psProduct)
		}
	}

	return res, nil
}

// Increase reserved quantities of products
func (ps Products) IncreaseReservedQTY(other Products) (Products, error) {
	if len(other) == 0 {
		return ps, nil
	}
	if len(other) > len(ps) {
		return nil, ErrNoProductItemInStock
	}

	result := Products{}
	for _, psProduct := range ps {
		reservedProduct := psProduct
		for _, otherProduct := range other {
			if !ps.Contain(otherProduct) {
				return nil, ErrNoProductItemInStock
			}
			if err := psProduct.IncreaseReservedQTY(otherProduct); err != nil {
				if err == ErrProductType {
					continue
				} else if err == ErrNotEnoughAvailableProducts {
					return nil, err
				}
			}
			reservedProduct = psProduct
		}
		result = append(result, reservedProduct)
	}

	return result, nil
}

// Decrease reserved quantities of products
func (ps Products) DecreaseReservedQTY(other Products) (Products, error) {
	if len(other) == 0 {
		return ps, nil
	}
	if len(other) > len(ps) {
		return nil, ErrNoProductItemInStock
	}

	result := Products{}
	for _, psProduct := range ps {
		reservedProduct := psProduct
		for _, otherProduct := range other {
			if !ps.Contain(otherProduct) {
				return nil, ErrNoProductItemInStock
			}
			if err := psProduct.DecreaseReservedQTY(otherProduct); err != nil {
				if err == ErrProductType {
					continue
				} else if err == ErrNotEnoughAvailableProducts {
					return nil, err
				}
			}
			reservedProduct = psProduct
		}
		result = append(result, reservedProduct)
	}

	if len(result) == 0 {
		return ps, nil
	}

	return result, nil
}

// Check products contain product
func (ps Products) Contain(product Product) bool {
	for _, productPs := range ps {
		if productPs.item.EqualTo(*product.item) {
			return true
		}
	}
	return false
}
