package mocks

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Save(ctx context.Context, w *models.Product) error {
	args := m.Called(ctx, w)
	return args.Error(0)
}

func (m *MockProductRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) GetAll(ctx context.Context) ([]*models.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Product), args.Error(1)
}

func (m *MockProductRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
