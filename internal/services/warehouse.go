package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"go.uber.org/zap"
)

const success = "success"

var (
	errEmptyParams = errors.New("empty parameters")
	errTransaction = errors.New("transaction execution error")
)

type Args struct {
	Warehouses []*models.Warehouse `json:"warehouses"`
}

type Reply struct {
	Message string `json:"message"`
}

type Service interface {
	SaveWarehouses(r *http.Request, args *Args, reply *Reply) (err error)
}

type service struct {
	warehouseRepository models.WarehouseRepository
	productRepository   models.ProductRepository
	logger              *zap.SugaredLogger
}

func NewService(warehouseRepo models.WarehouseRepository, productRepo models.ProductRepository, logger *zap.SugaredLogger) *service {
	return &service{
		warehouseRepository: warehouseRepo,
		productRepository:   productRepo,
		logger:              logger,
	}
}

func (s *service) SaveWarehouses(r *http.Request, args *Args, reply *Reply) (err error) {
	warehouses := args.Warehouses
	if len(warehouses) == 0 {
		err = errEmptyParams
		s.logger.Info(err)
		return
	}

	for _, warehouse := range warehouses {
		if err = warehouse.Validate(); err != nil {
			s.logger.Info(err)
			return err
		}
	}

	var queryErr error
	err = s.warehouseRepository.InTransaction(context.Background(), func(ctx context.Context) error {
		for _, warehouse := range warehouses {
			if queryErr = s.warehouseRepository.Save(ctx, warehouse); queryErr != nil {
				return queryErr
			}
		}
		return nil
	})

	switch {
	case queryErr != nil:
		s.logger.Info(queryErr)
		return errTransaction
	case err != nil:
		s.logger.Info(err)
		return errTransaction
	}

	s.logger.Info(success)
	reply.Message = "warehouse(s) was successfully added"
	return nil
}
