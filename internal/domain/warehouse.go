package domain

import (
	"errors"
)

var (
	ErrWarehouseIsNotAvailable = errors.New("warehouse is not available")
	ErrEmptyWarehouseName      = errors.New("warehouse name cannot be empty")
	ErrEmptyProducts           = errors.New("insufficient number of products")
)

// Warehouse
type Warehouse struct {
	id          uint
	name        string
	isAvailable bool
	products    Products
}

// Create new warehouse
func NewWarehouse(name string, isAvailable bool, products Products) (Warehouse, error) {
	if name == "" {
		return Warehouse{}, ErrEmptyWarehouseName
	}
	return Warehouse{
		name:        name,
		isAvailable: isAvailable,
		products:    products,
	}, nil
}

// Get warehouse id
func (w Warehouse) GetId() uint {
	return w.id
}

// Set warehouse item id
func (w *Warehouse) SetId(id uint) {
	w.id = id
}

// Get warehouse name
func (w Warehouse) GetName() string {
	return w.name
}

// Set warehouse name
func (w *Warehouse) SetName(name string) {
	w.name = name
}

// Get warehouse availability flag
func (w Warehouse) GetIsAvailable() bool {
	return w.isAvailable
}

// Set warehouse availability flag
func (w *Warehouse) SetIsAvailable(isAvailable bool) {
	w.isAvailable = isAvailable
}

// Get warehouse products
func (w Warehouse) GetProducts() Products {
	return w.products
}

// Set warehouse products
func (w *Warehouse) SetProducts(products Products) {
	w.products = products
}

// Check warehouse is available
func (w Warehouse) IsAvailable() bool {
	return w.isAvailable
}

// Reserve products from the warehouse
func (w *Warehouse) Reserve(products Products) error {
	if !w.IsAvailable() {
		return ErrWarehouseIsNotAvailable
	}
	products, err := w.products.IncreaseReservedQTY(products)
	if err != nil {
		return err
	}
	w.products = products

	return nil
}

// Remove the reserve from products in the warehouse
func (w *Warehouse) Release(products Products) error {
	if !w.IsAvailable() {
		return ErrWarehouseIsNotAvailable
	}
	products, err := w.products.DecreaseReservedQTY(products)
	if err != nil {
		return err
	}
	w.products = products

	return nil
}
