package domain

import "github.com/gofrs/uuid"

type DiningTableStatus string

const (
	DiningTableStatusAvailable DiningTableStatus = "available"
	DiningTableStatusOccupied  DiningTableStatus = "active"
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
