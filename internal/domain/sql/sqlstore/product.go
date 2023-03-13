package sqlstore

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
)

type ProductDto struct {
	Id           uint      `db:"id"`
	Code         uuid.UUID `db:"code"`
	Name         string    `db:"name"`
	Size         uint      `db:"size"`
	AvailableQTY uint      `db:"available_qty"`
	ReservedQTY  uint      `db:"reserved_qty"`
}

// Create product from dto
func (d ProductDto) ToAggregate() domain.Product {
	item := domain.Item{
		Id:   d.Id,
		Code: d.Code,
		Name: d.Name,
		Size: d.Size,
	}
	product := domain.Product{}
	product.SetItem(&item)
	product.SetAvailableQTY(d.AvailableQTY)
	product.SetReservedQTY(d.ReservedQTY)

	return product
}

// Create dto from product
func FromProduct(p domain.Product) ProductDto {
	return ProductDto{
		Id:           p.GetId(),
		Code:         p.GetCode(),
		Name:         p.GetName(),
		Size:         p.GetSize(),
		AvailableQTY: p.GetAvailableQTY(),
		ReservedQTY:  p.GetReservedQTY(),
	}
}

// Save product
func SaveProduct(ctx context.Context, db SqlxDatabase, dto ProductDto) error {
	query, _, err := prodPreparedDs.Update().
		Set(goqu.Record{
			ProductNameColumn: "$1",
			ProductSizeColumn: "$2",
		}).Where(goqu.And(prodIdCol.Eq("$3"), prodCodeCol.Eq("4"))).ToSQL()
	if err != nil {
		return err
	}
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}
	var lastId uint
	stmt.GetContext(ctx, &lastId, dto.Name, dto.Size, dto.Id, dto.Code)
	return nil
}

// Get the quantity of product in all warehouses by id
func GetProductTotalByItemId(ctx context.Context, db SqlxDatabase, id uint) (domain.Product, error) {
	query, _, err := prodPreparedDs.Select(
		prodIdCol, prodCodeCol, prodNameCol, prodSizeCol,
		goqu.SUM(plcQTYCol).As("available_qty"),
		goqu.COALESCE(goqu.SUM(rsvQTYCol), "$1").As("reserved_qty")).
		Join(plcT, goqu.On(prodIdCol.Eq(plcProductIdCol))).
		LeftJoin(rsvT, goqu.On(plcIdCol.Eq(rsvPlacementIdCol))).
		Where(prodIdCol.Eq("$2")).
		GroupBy(prodIdCol).
		ToSQL()
	if err != nil {
		return domain.Product{}, err
	}
	dto := new(ProductDto)
	err = db.GetContext(ctx, dto, query, 0, id)
	if err != nil {
		return domain.Product{}, domain.ErrProductNotFound
	}
	return dto.ToAggregate(), nil
}

// Get the quantity of product in all warehouses by code
func GetProductTotalByItemCode(ctx context.Context, db SqlxDatabase, code uuid.UUID) (domain.Product, error) {
	query, _, err := prodPreparedDs.Select(
		prodIdCol, prodCodeCol, prodNameCol, prodSizeCol,
		goqu.SUM(plcQTYCol).As("available_qty"),
		goqu.COALESCE(goqu.SUM(rsvQTYCol), "$1").As("reserved_qty")).
		Join(plcT, goqu.On(prodIdCol.Eq(plcProductIdCol))).
		LeftJoin(rsvT, goqu.On(plcIdCol.Eq(rsvPlacementIdCol))).
		Where(prodCodeCol.Eq("$2")).
		GroupBy(prodIdCol).
		ToSQL()
	if err != nil {
		return domain.Product{}, err
	}
	dto := new(ProductDto)
	err = db.GetContext(ctx, dto, query, 0, code)
	if err != nil {
		return domain.Product{}, domain.ErrProductNotFound
	}
	return dto.ToAggregate(), nil
}

// Get the quantity of product in warehouse by id and warehouse id
func GetProductInStockByItemId(ctx context.Context, db SqlxDatabase, id uint, warehouse_id uint) (domain.Product, error) {
	query, _, err := prodPreparedDs.Select(
		prodIdCol, prodCodeCol, prodNameCol, prodSizeCol,
		plcQTYCol.As("available_qty"),
		goqu.COALESCE(rsvQTYCol, "$1").As("reserved_qty")).
		Join(plcT, goqu.On(prodIdCol.Eq(plcProductIdCol))).
		LeftJoin(rsvT, goqu.On(plcIdCol.Eq(rsvPlacementIdCol))).
		Where(goqu.And(prodIdCol.Eq("$2"), plcWarehouseIdCol.Eq("$3"))).
		ToSQL()
	if err != nil {
		return domain.Product{}, err
	}
	dto := new(ProductDto)
	err = db.GetContext(ctx, dto, query, 0, id, warehouse_id)
	if err != nil {
		return domain.Product{}, domain.ErrProductNotFound
	}
	return dto.ToAggregate(), nil
}

// Get the quantity of product in warehouse by code and warehouse id
func GetProductInStockByItemCode(ctx context.Context, db SqlxDatabase, code uuid.UUID, warehouse_id uint) (domain.Product, error) {
	query, _, err := prodPreparedDs.Select(
		prodIdCol, prodCodeCol, prodNameCol, prodSizeCol,
		plcQTYCol.As("available_qty"),
		goqu.COALESCE(rsvQTYCol, "$1").As("reserved_qty")).
		Join(plcT, goqu.On(prodIdCol.Eq(plcProductIdCol))).
		LeftJoin(rsvT, goqu.On(plcIdCol.Eq(rsvPlacementIdCol))).
		Where(goqu.And(prodCodeCol.Eq("$2"), plcWarehouseIdCol.Eq("$3"))).
		ToSQL()
	if err != nil {
		return domain.Product{}, err
	}
	dto := new(ProductDto)
	err = db.GetContext(ctx, dto, query, 0, code, warehouse_id)
	if err != nil {
		return domain.Product{}, domain.ErrProductNotFound
	}
	return dto.ToAggregate(), nil
}

// Get the quantity of all products in warehouse by warehouse id
func GetAllProductsInStockByWarehouseId(ctx context.Context, db SqlxDatabase, warehouse_id uint) (domain.Products, error) {
	query, _, err := prodPreparedDs.Select(
		prodIdCol, prodCodeCol, prodNameCol, prodSizeCol,
		plcQTYCol.As("available_qty"),
		goqu.COALESCE(rsvQTYCol, "$1").As("reserved_qty")).
		Join(plcT, goqu.On(prodIdCol.Eq(plcProductIdCol))).
		LeftJoin(rsvT, goqu.On(plcIdCol.Eq(rsvPlacementIdCol))).
		Where(plcWarehouseIdCol.Eq("$2")).
		ToSQL()
	if err != nil {
		return domain.Products{}, err
	}
	var dtos []ProductDto
	err = db.SelectContext(ctx, &dtos, query, 0, warehouse_id)
	if err != nil {
		return domain.Products{}, domain.ErrProductNotFound
	}
	products := make(domain.Products, len(dtos))
	for index, dto := range dtos {
		products[index] = dto.ToAggregate()
	}
	return products, nil
}
