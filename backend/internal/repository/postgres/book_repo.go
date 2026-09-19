package postgres

import (
	"context"

	"github.com/azrilpramudia/libre-go/internal/domain/book"
	"github.com/azrilpramudia/libre-go/internal/repository/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type bookRepo struct {
	q *sqlc.Queries
}

func NewBookRepository(q *sqlc.Queries) book.Repository {
	return &bookRepo{q: q}
}

func (r *bookRepo) Create(ctx context.Context, b *book.Book) (*book.Book, error) {
	row, err := r.q.CreateBook(ctx, sqlc.CreateBookParams{
		Title:     b.Title,
		Isbn:      pgtype.Text{String: b.ISBN, Valid: b.ISBN != ""},
		Stock:     b.Stock,
		Available: b.Stock,
	})
	if err != nil {
		return nil, err
	}
	return toDomainBook(row), nil
}

func (r *bookRepo) GetByID(ctx context.Context, id uuid.UUID) (*book.Book, error) {
	row, err := r.q.GetBookByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return nil, err
	}
	return toDomainBook(row), nil
}

func (r *bookRepo) List(ctx context.Context, limit, offset int32) ([]*book.Book, error) {
	rows, err := r.q.ListBooks(ctx, sqlc.ListBooksParams{Limit: limit, Offset: offset})
	if err != nil {
		return nil, err
	}
	result := make([]*book.Book, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainBook(row))
	}
	return result, nil
}

func (r *bookRepo) Update(ctx context.Context, b *book.Book) (*book.Book, error) {
	return nil, nil
}

func (r *bookRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBook(ctx, pgtype.UUID{Bytes: id, Valid: true})
}

func toDomainBook(row sqlc.Book) *book.Book {
	return &book.Book{
		ID:        row.ID.Bytes,
		Title:     row.Title,
		ISBN:      row.Isbn.String,
		Stock:     row.Stock,
		Available: row.Available,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}