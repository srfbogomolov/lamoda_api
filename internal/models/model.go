package models

import "errors"

var (
	errEmpty        = errors.New("cannot be empty")
	errLessThanZero = errors.New("cannot be less than zero")
)

type Model interface {
	Validate() error
}
