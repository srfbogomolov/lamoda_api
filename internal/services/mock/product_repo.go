package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetTotalByItemId(ctx context.Context, id uint) (domain.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) Save(ctx context.Context, product domain.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) GetTotalByItemCode(ctx context.Context, code uuid.UUID) (domain.Product, error) {
	args := m.Called(ctx, code)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) GetInStockByItemId(ctx context.Context, id uint, warehouse_id uint) (domain.Product, error) {
	args := m.Called(ctx, id, warehouse_id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) GetInStockByItemCode(ctx context.Context, code uuid.UUID, warehouse_id uint) (domain.Product, error) {
	args := m.Called(ctx, code, warehouse_id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) GetAllInStockByWarehouseId(ctx context.Context, warehouse_id uint) (domain.Products, error) {
	args := m.Called(ctx, warehouse_id)
	return args.Get(0).(domain.Products), args.Error(1)
}

func (m *MockProductRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
