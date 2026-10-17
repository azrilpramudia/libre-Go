package loan

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID         uuid.UUID
	BookID     uuid.UUID
	MemberID   uuid.UUID
	BorrowedAt time.Time
	DueAt      time.Time
	ReturnedAt *time.Time
	Status     string
	FineAmount int
}
