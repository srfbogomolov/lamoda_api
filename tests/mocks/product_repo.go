package mocks

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Save(ctx context.Context, w *models.Product) (string, error) {
	args := m.Called(ctx, w)
	return args.String(0), args.Error(1)
}

func (m *MockProductRepository) FindByCode(ctx context.Context, code string) (*models.Product, error) {
	args := m.Called(ctx, code)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Find(ctx context.Context) ([]*models.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Product), args.Error(1)
}

func (m *MockProductRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
