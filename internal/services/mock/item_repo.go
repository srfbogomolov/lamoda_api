package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockItemRepository struct {
	mock.Mock
}

func (m *MockItemRepository) GetById(ctx context.Context, id uint) (*domain.Item, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemRepository) GetByCode(ctx context.Context, code uuid.UUID) (*domain.Item, error) {
	args := m.Called(ctx, code)
	return args.Get(0).(*domain.Item), args.Error(1)
}

func (m *MockItemRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
