package models

import "errors"

var (
	ErrEmpty           = errors.New("cannot be empty")
	ErrLessZero        = errors.New("cannot be less than zero")
	ErrLessOrEqualZero = errors.New("cannot be less or equal zero")
)

type Model interface {
	Validate() error
}
