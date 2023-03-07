package mocks

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockWarehouseRepository struct {
	mock.Mock
}

func (m *MockWarehouseRepository) Save(ctx context.Context, w *models.Warehouse) (int, error) {
	args := m.Called(ctx, w)
	return args.Int(0), args.Error(1)
}

func (m *MockWarehouseRepository) FindById(ctx context.Context, id int) (*models.Warehouse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) Find(ctx context.Context) ([]*models.Warehouse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
