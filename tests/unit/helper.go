package unit_test

import "errors"

var (
	errEmpty        = errors.New("cannot be empty")
	errLessThanZero = errors.New("cannot be less than zero")
)
