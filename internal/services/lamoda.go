package services

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/srfbogomolov/warehouse_api/internal/domain/sql"
)

const (
	StatusOK = "OK"
)

var (
	ErrEmptyWarehouseArg    = errors.New("warehouse id is not specified")
	ErrEmptyProductCodesArg = errors.New("product codes are not specified")
	ErrServer               = errors.New("server error")
	ErrTransaction          = errors.New("transaction execution error")
)

type GetInStockArgs struct {
	WarehouseId uint `json:"warehouse_id"`
}

type AvailableProduct struct {
	Id           uint      `json:"id"`
	Code         uuid.UUID `json:"code"`
	Name         string    `json:"name"`
	Size         uint      `json:"size"`
	AvailableQTY uint      `json:"available_qty"`
}

type GetInStockReply struct {
	AvailableProducts []AvailableProduct `json:"available_products"`
	Status            string             `json:"status"`
}

type ReleaseReserveArgs struct {
	WarehouseId  uint        `json:"warehouse_id"`
	ProductCodes []uuid.UUID `json:"product_codes"`
}

type ReleaseReserveReply struct {
	Status string `json:"status"`
}

type Service interface {
	GetInStock(r *http.Request, args *GetInStockArgs, reply *GetInStockReply) error
	Reserve(r *http.Request, args *ReleaseReserveArgs, reply *ReleaseReserveReply) error
	Release(r *http.Request, args *ReleaseReserveArgs, reply *ReleaseReserveReply) error
}

type LamodaService struct {
	Item      domain.ItemRepository
	Product   domain.ProductRepository
	Warehouse domain.WarehouseRepository
}

func NewLamodaService(db *sqlx.DB) *LamodaService {
	return &LamodaService{
		Item:      sql.NewSqlItemRepository(db),
		Product:   sql.NewSqlProductRepository(db),
		Warehouse: sql.NewSqlWarehouseRepository(db),
	}
}

// Make a map of the code-quantity type
func CountProductCodes(codes []uuid.UUID) map[uuid.UUID]uint {
	counter := make(map[uuid.UUID]uint)
	for _, code := range codes {
		counter[code] = counter[code] + 1
	}
	return counter
}

// Get products by their code
func GetProductsByCodes(ls *LamodaService, codes []uuid.UUID) (domain.Products, error) {
	products := domain.Products{}
	for code, qty := range CountProductCodes(codes) {
		item, err := ls.Item.GetByCode(context.Background(), code)
		if err != nil {
			return domain.Products{}, err
		}
		product := domain.Product{}
		product.SetItem(item)
		product.SetReservedQTY(qty)
		products = append(products, product)
	}
	return products, nil
}

// Reserve products in a warehouse for delivery
func (ls *LamodaService) Reserve(r *http.Request, args *ReleaseReserveArgs, reply *ReleaseReserveReply) error {
	warehouseId := args.WarehouseId
	productCodes := args.ProductCodes
	if warehouseId == 0 {
		return ErrEmptyWarehouseArg
	} else if len(productCodes) == 0 {
		return ErrEmptyProductCodesArg
	}

	warehouse, err := ls.Warehouse.GetById(context.Background(), warehouseId)
	if err != nil {
		return err
	}
	if !warehouse.IsAvailable() {
		return domain.ErrWarehouseIsNotAvailable
	}

	products, err := GetProductsByCodes(ls, productCodes)
	if err != nil {
		return err
	}

	if err := warehouse.Reserve(products); err != nil {
		return err
	}

	var queryErr error
	transactionErr := ls.Warehouse.InTransaction(context.Background(), func(ctx context.Context) error {
		if queryErr = ls.Warehouse.Save(ctx, warehouse); err != nil {
			return queryErr
		}
		return nil
	})

	if transactionErr != nil {
		log.Println(transactionErr)
		return ErrTransaction
	} else if queryErr != nil {
		return queryErr
	}

	reply.Status = StatusOK

	return nil
}

// Release products reserve
func (ls *LamodaService) Release(r *http.Request, args *ReleaseReserveArgs, reply *ReleaseReserveReply) error {
	warehouseId := args.WarehouseId
	productCodes := args.ProductCodes
	if warehouseId == 0 {
		return ErrEmptyWarehouseArg
	} else if len(productCodes) == 0 {
		return ErrEmptyProductCodesArg
	}

	warehouse, err := ls.Warehouse.GetById(context.Background(), warehouseId)
	if err != nil {
		return err
	}
	if !warehouse.IsAvailable() {
		return domain.ErrWarehouseIsNotAvailable
	}

	products, err := GetProductsByCodes(ls, productCodes)
	if err != nil {
		return err
	}

	if err := warehouse.Release(products); err != nil {
		return err
	}

	var queryErr error
	transactionErr := ls.Warehouse.InTransaction(context.Background(), func(ctx context.Context) error {
		if queryErr = ls.Warehouse.Save(ctx, warehouse); err != nil {
			return queryErr
		}
		return nil
	})

	if transactionErr != nil {
		log.Println(transactionErr)
		return ErrTransaction
	} else if queryErr != nil {
		return queryErr
	}

	reply.Status = StatusOK

	return nil
}

// Get the number of remaining products in warehouse
func (ls *LamodaService) GetInStock(r *http.Request, args *GetInStockArgs, reply *GetInStockReply) error {
	warehouseId := args.WarehouseId
	if warehouseId == 0 {
		return ErrEmptyWarehouseArg
	}

	warehouse, err := ls.Warehouse.GetById(context.Background(), warehouseId)
	if err != nil {
		return err
	}
	if !warehouse.IsAvailable() {
		return domain.ErrWarehouseIsNotAvailable
	}

	productsInStock := warehouse.GetProducts()
	for _, product := range productsInStock {
		if product.GetAvailableQTY() == 0 {
			continue
		}
		availableProduct := AvailableProduct{}
		availableProduct.Id = product.GetId()
		availableProduct.Code = product.GetCode()
		availableProduct.Name = product.GetName()
		availableProduct.Size = product.GetSize()
		availableProduct.AvailableQTY = product.GetAvailableQTY()
		reply.AvailableProducts = append(reply.AvailableProducts, availableProduct)
	}

	reply.Status = StatusOK

	return nil
}
