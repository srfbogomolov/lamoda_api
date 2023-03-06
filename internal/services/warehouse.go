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
	Products   []*models.Product   `json:"products"`
}

type Reply struct {
	Message string `json:"message"`
}

type Service interface {
	SaveWarehouses(r *http.Request, args *Args, reply *Reply) (err error)
	SaveProducts(r *http.Request, args *Args, reply *Reply) (err error)
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

func (s *service) SaveProducts(r *http.Request, args *Args, reply *Reply) (err error) {
	products := args.Products
	if len(products) == 0 {
		err = errEmptyParams
		s.logger.Info(err)
		return
	}

	for _, product := range products {
		if err = product.Validate(); err != nil {
			s.logger.Info(err)
			return err
		}
	}

	var queryErr error
	err = s.productRepository.InTransaction(context.Background(), func(ctx context.Context) error {
		for _, product := range products {
			if queryErr = s.productRepository.Save(ctx, product); queryErr != nil {
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
	reply.Message = "product(s) was successfully added"
	return nil
}
