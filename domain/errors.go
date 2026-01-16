package domain

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrConflict   = errors.New("resource conflict")
	ErrUnexpected = errors.New("unexpected error")
)
