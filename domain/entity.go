package domain

import "github.com/gofrs/uuid"

type DiningTableStatus string

func (d DiningTableStatus) String() string {
	return string(d)
}

const (
	DiningTableStatusAvailable DiningTableStatus = "available"
	DiningTableStatusActive    DiningTableStatus = "active"
)

type DiningTable struct {
	ID          uuid.UUID
	TableNumber int
	Capacity    int
	TableStatus DiningTableStatus
}

type Owner struct {
	ID    uuid.UUID
	Name  string
	Email string
}
