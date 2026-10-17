package service

import (
	"context"
	"time"

	"github.com/azrilpramudia/libre-go/internal/domain/loan"
	"github.com/google/uuid"
)

type LoanService struct {
	repo loan.Repository
}

func NewLoanService(repo loan.Repository) *LoanService {
	return &LoanService{repo: repo}
}

const defaultLoanDuration = 7 * 24 * time.Hour

func (s *LoanService) BorrowBook(ctx context.Context, bookID, memberID uuid.UUID) (*loan.Loan, error) {
	dueAt := time.Now().Add(defaultLoanDuration)
	return s.repo.Borrow(ctx, bookID, memberID, dueAt)
}

func (s *LoanService) ReturnBook(ctx context.Context, loanID uuid.UUID) (*loan.Loan, error) {
	return s.repo.Return(ctx, loanID)
}

func (s *LoanService) ListMemberLoans(ctx context.Context, memberID uuid.UUID) ([]*loan.Loan, error) {
	return s.repo.ListByMember(ctx, memberID)
}
