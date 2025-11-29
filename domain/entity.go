package domain

import "github.com/gofrs/uuid"

type Owner struct {
	ID    uuid.UUID
	Name  string
	Email string
}
