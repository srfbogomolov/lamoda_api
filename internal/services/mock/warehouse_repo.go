package mock

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockWarehouseRepository struct {
	mock.Mock
}

func (m *MockWarehouseRepository) Save(ctx context.Context, warehouse domain.Warehouse) error {
	args := m.Called(ctx, warehouse)
	return args.Error(0)
}

func (m *MockWarehouseRepository) GetById(ctx context.Context, id uint) (domain.Warehouse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
