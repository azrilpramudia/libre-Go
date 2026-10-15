package postgres

import (
	"context"
	"errors"

	"github.com/azrilpramudia/libre-go/internal/domain/author"
	"github.com/azrilpramudia/libre-go/internal/repository/postgres/sqlc"
	"github.com/jackc/pgx/v5"
)

type authorRepo struct {
	q *sqlc.Queries
}

func NewAuthorRepository(q *sqlc.Queries) author.Repository {
	return &authorRepo{q: q}
}

func (r *authorRepo) FindOrCreateByName(ctx context.Context, name string) (*author.Author, error) {
	row, err := r.q.GetAuthorByName(ctx, name)
	if err == nil {
		return &author.Author{ID: row.ID.Bytes, Name: row.Name}, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	created, err := r.q.CreateAuthor(ctx, name)
	if err != nil {
		return nil, err
	}
	return &author.Author{ID: created.ID.Bytes, Name: created.Name}, nil
}
