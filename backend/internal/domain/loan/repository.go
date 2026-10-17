package loan

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrBookNotAvailable    = errors.New("book not available")
	ErrLoanAlreadyReturned = errors.New("loan already returned")
)

type Repository interface {
	Borrow(ctx context.Context, bookID, memberID uuid.UUID, dueAt time.Time) (*Loan, error)
	Return(ctx context.Context, loanID uuid.UUID) (*Loan, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Loan, error)
	ListByMember(ctx context.Context, memberID uuid.UUID) ([]*Loan, error)
}
