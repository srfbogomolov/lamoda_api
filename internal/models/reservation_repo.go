package models

import "context"

type ReservationRepository interface {
	Save(ctx context.Context, reservation *Reservation) (uint, error)
	Find(ctx context.Context) ([]*Reservation, error)
	FindById(ctx context.Context, id uint) (*Reservation, error)
	FindByPlacementId(ctx context.Context, placementId uint) (*Reservation, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}
