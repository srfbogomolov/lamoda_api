package services

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/srfbogomolov/warehouse_api/internal/models"
	"go.uber.org/zap"
)

const success = "success"

var (
	errEmptyParams             = errors.New("empty parameters")
	errTransaction             = errors.New("transaction execution error")
	errWarehouseIsNotAvailable = errors.New("warehouse is not available")
	errNotEnoughProducts       = errors.New("not enough products")
)

type Args struct {
	Warehouses []*models.Warehouse `json:"warehouses"`
	Products   []*models.Product   `json:"products"`
	Placements []*models.Placement `json:"placements"`
}

type Reply struct {
	Message string `json:"message"`
}

type Service interface {
	SaveWarehouses(r *http.Request, args *Args, reply *Reply) (err error)
	SaveProducts(r *http.Request, args *Args, reply *Reply) (err error)
	SavePlacements(r *http.Request, args *Args, reply *Reply) (err error)
}

type service struct {
	warehouseRepository models.WarehouseRepository
	productRepository   models.ProductRepository
	placementRepository models.PlacementRepository
	logger              *zap.SugaredLogger
}

func NewService(wRepo models.WarehouseRepository, pRepo models.ProductRepository, plRepo models.PlacementRepository, logger *zap.SugaredLogger) *service {
	return &service{
		warehouseRepository: wRepo,
		productRepository:   pRepo,
		placementRepository: plRepo,
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
			if _, queryErr = s.warehouseRepository.Save(ctx, warehouse); queryErr != nil {
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
	reply.Message = "warehouse(s) was(were) successfully added"
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
			if _, queryErr = s.productRepository.Save(ctx, product); queryErr != nil {
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
	reply.Message = "product(s) was(were) successfully added"
	return nil
}

func (s *service) SavePlacements(r *http.Request, args *Args, reply *Reply) (err error) {
	placements := args.Placements
	if len(placements) == 0 {
		err = errEmptyParams
		s.logger.Info(err)
		return err
	}

	for _, placement := range placements {
		if err = placement.Validate(); err != nil {
			s.logger.Info(err)
			return err
		}
	}

	var queryErr error
	var logicErr error
	err = s.placementRepository.InTransaction(context.Background(), func(ctx context.Context) error {
		for _, placement := range placements {
			warehouse, queryErr := s.warehouseRepository.FindById(ctx, placement.WarehouseId)
			if queryErr != nil {
				return queryErr
			}
			if !warehouse.IsAvailable {
				logicErr = errWarehouseIsNotAvailable
				return errWarehouseIsNotAvailable
			}

			qtySum, queryErr := s.placementRepository.FindQTYSumByProductCode(ctx, placement.ProductCode)
			if queryErr != nil && queryErr != sql.ErrNoRows {
				return queryErr
			}

			product, queryErr := s.productRepository.FindByCode(ctx, placement.ProductCode)
			if queryErr != nil {
				return queryErr
			}
			if product.QTY < placement.QTY+qtySum {
				logicErr = errNotEnoughProducts
				return logicErr
			}

			if _, queryErr = s.placementRepository.Save(ctx, placement); queryErr != nil {
				return queryErr
			}
		}
		return nil
	})

	switch {
	case queryErr != nil:
		s.logger.Info(queryErr)
		return errTransaction
	case logicErr != nil:
		s.logger.Info(logicErr)
		return logicErr
	case err != nil:
		s.logger.Info(err)
		return errTransaction
	}

	s.logger.Info(success)
	reply.Message = "placement(s) was(were) successfully added"
	return nil
}
