package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrItemNotFound = errors.New("item not found")
)

type ItemRepository interface {
	GetById(ctx context.Context, id uint) (*Item, error)
	GetByCode(ctx context.Context, code uuid.UUID) (*Item, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}
