package models

import (
	"errors"
)

// todo удалить
var (
	ErrEmpty           = errors.New("cannot be empty")
	ErrLessZero        = errors.New("cannot be less than zero")
	ErrLessOrEqualZero = errors.New("cannot be less or equal zero")
	ErrIncorrectUUID   = errors.New("cannot be incorrect")
)
