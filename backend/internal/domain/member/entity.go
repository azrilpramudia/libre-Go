package member

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID           uuid.UUID
	FullName     string
	Email        string
	Phone        string
	PasswordHash string
	CreatedAt    time.Time
}
