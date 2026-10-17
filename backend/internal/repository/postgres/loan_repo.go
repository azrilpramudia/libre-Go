package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/azrilpramudia/libre-go/internal/domain/loan"
	"github.com/azrilpramudia/libre-go/internal/repository/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type loanRepo struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func NewLoanRepository(pool *pgxpool.Pool, q *sqlc.Queries) loan.Repository {
	return &loanRepo{pool: pool, q: q}
}

func (r *loanRepo) Borrow(ctx context.Context, bookID, memberID uuid.UUID, dueAt time.Time) (*loan.Loan, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	_, err = qtx.DecrementBookAvailable(ctx, pgtype.UUID{Bytes: bookID, Valid: true})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, loan.ErrBookNotAvailable
	}
	if err != nil {
		return nil, err
	}

	row, err := qtx.CreateLoan(ctx, sqlc.CreateLoanParams{
		BookID:   pgtype.UUID{Bytes: bookID, Valid: true},
		MemberID: pgtype.UUID{Bytes: memberID, Valid: true},
		DueAt:    pgtype.Timestamptz{Time: dueAt, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return toDomainLoan(row), nil
}

func (r *loanRepo) Return(ctx context.Context, loanID uuid.UUID) (*loan.Loan, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	existing, err := qtx.GetLoanByID(ctx, pgtype.UUID{Bytes: loanID, Valid: true})
	if err != nil {
		return nil, err
	}

	fine := calculateFine(existing.DueAt.Time, time.Now())

	updated, err := qtx.ReturnLoan(ctx, sqlc.ReturnLoanParams{
		ID:         pgtype.UUID{Bytes: loanID, Valid: true},
		FineAmount: int32(fine),
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, loan.ErrLoanAlreadyReturned
	}
	if err != nil {
		return nil, err
	}

	if _, err := qtx.IncrementBookAvailable(ctx, existing.BookID); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return toDomainLoan(updated), nil
}

func (r *loanRepo) GetByID(ctx context.Context, id uuid.UUID) (*loan.Loan, error) {
	row, err := r.q.GetLoanByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return nil, err
	}
	return toDomainLoan(row), nil
}

func (r *loanRepo) ListByMember(ctx context.Context, memberID uuid.UUID) ([]*loan.Loan, error) {
	rows, err := r.q.ListLoansByMember(ctx, pgtype.UUID{Bytes: memberID, Valid: true})
	if err != nil {
		return nil, err
	}
	result := make([]*loan.Loan, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainLoan(row))
	}
	return result, nil
}

func calculateFine(dueAt, returnedAt time.Time) int {
	const finePerDay = 2000
	if returnedAt.Before(dueAt) || returnedAt.Equal(dueAt) {
		return 0
	}
	daysLate := int(returnedAt.Sub(dueAt).Hours() / 24)
	if daysLate < 1 {
		daysLate = 1
	}
	return daysLate * finePerDay
}

func toDomainLoan(row sqlc.Loan) *loan.Loan {
	l := &loan.Loan{
		ID:         row.ID.Bytes,
		BookID:     row.BookID.Bytes,
		MemberID:   row.MemberID.Bytes,
		BorrowedAt: row.BorrowedAt.Time,
		DueAt:      row.DueAt.Time,
		Status:     row.Status,
	}
	if row.ReturnedAt.Valid {
		l.ReturnedAt = &row.ReturnedAt.Time
	}
	return l
}
