package book

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, b *Book) (*Book, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Book, error)
	List(ctx context.Context, limit, offset int32) ([]*Book, error)
	Update(ctx context.Context, b *Book) (*Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
}