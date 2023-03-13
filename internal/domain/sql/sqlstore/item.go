package sqlstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
)

type ItemDto struct {
	Id   uint      `db:"id" json:"id"`
	Code uuid.UUID `db:"code" json:"code"`
	Name string    `db:"name" json:"name"`
	Size uint      `db:"size" json:"size"`
}

// Get item by id
func GetItemById(ctx context.Context, db SqlxDatabase, id uint) (*domain.Item, error) {
	query, _, err := prodPreparedDs.Where(prodIdCol.Eq("$1")).ToSQL()
	if err != nil {
		return &domain.Item{}, err
	}
	dto := new(ItemDto)
	err = db.GetContext(ctx, dto, query, id)
	if err != nil {
		return &domain.Item{}, domain.ErrItemNotFound
	}
	return &domain.Item{
		Id:   dto.Id,
		Code: dto.Code,
		Name: dto.Name,
		Size: dto.Size,
	}, nil
}

// Get item by code
func GetItemByCode(ctx context.Context, db SqlxDatabase, code uuid.UUID) (*domain.Item, error) {
	query, _, err := prodPreparedDs.Where(prodCodeCol.Eq("$1")).ToSQL()
	if err != nil {
		return &domain.Item{}, err
	}
	dto := new(ItemDto)
	err = db.GetContext(ctx, dto, query, code)
	if err != nil {
		return &domain.Item{}, domain.ErrItemNotFound
	}
	return &domain.Item{
		Id:   dto.Id,
		Code: dto.Code,
		Name: dto.Name,
		Size: dto.Size,
	}, nil
}
