package domain

import (
	"errors"
)

var (
	ErrConflict   = errors.New("resource conflict")
	ErrUnexpected = errors.New("unexpected error")
)
