package models

type Reservation struct {
	Id          uint `db:"id" json:"id" validate:"required,gt=0"`
	PlacementId uint `db:"placement_id" json:"placement_id" validate:"required,gt=0"`
	QTY         uint `db:"qty" json:"qty" validate:"required,gt=0"`
}

func NewReservation(id uint, placementId uint, qty uint) *Reservation {
	return &Reservation{
		Id:          id,
		PlacementId: placementId,
		QTY:         qty,
	}
}
