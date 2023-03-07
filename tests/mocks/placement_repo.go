package mocks

import (
	"context"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockPlacementRepository struct {
	mock.Mock
}

func (m *MockPlacementRepository) Save(ctx context.Context, w *models.Placement) (int, error) {
	args := m.Called(ctx, w)
	return args.Int(0), args.Error(1)
}

func (m *MockPlacementRepository) FindById(ctx context.Context, id int) (*models.Placement, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Placement), args.Error(1)
}

func (m *MockPlacementRepository) Find(ctx context.Context) ([]*models.Placement, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Placement), args.Error(1)
}

func (m *MockPlacementRepository) FindQTYSumByProductCode(ctx context.Context, productCode string) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *MockPlacementRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
