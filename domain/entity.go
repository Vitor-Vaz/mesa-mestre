package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Owner struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
}
