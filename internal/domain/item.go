package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmptyItemName = errors.New("item name cannot be empty")
)

// Item
type Item struct {
	Id   uint
	Code uuid.UUID
	Name string
	Size uint
}

// Create new item
func NewItem(name string, size uint) (*Item, error) {
	if name == "" {
		return &Item{}, ErrEmptyItemName
	}
	return &Item{
		Id:   0,
		Code: uuid.New(),
		Name: name,
		Size: size,
	}, nil
}

// Check item is equal to another item
func (i Item) EqualTo(other Item) bool {
	return other.Id == i.Id && other.Code == i.Code
}
