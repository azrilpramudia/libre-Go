package member

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID        uuid.UUID
	FullName  string
	Email     string
	Phone     string
	CreatedAt time.Time
}
