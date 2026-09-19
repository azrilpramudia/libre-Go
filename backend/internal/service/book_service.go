package service

import (
	"context"

	"github.com/azrilpramudia/libre-go/internal/domain/book"
	"github.com/google/uuid"
)

type BookService struct {
	repo book.Repository
}

func NewBookRepository(repo book.Repository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, b *book.Book) (*book.Book, error) {
	if b.Stock < 1 {
		b.Stock = 1
	}
	return s.repo.Create(ctx, b)
}

func (s *BookService) GetBook(ctx context.Context, id uuid.UUID) (*book.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *BookService) ListBooks(ctx context.Context, limit, offset int32) ([]*book.Book, error) {
	return s.repo.List(ctx, limit, offset)
}