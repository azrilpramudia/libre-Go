package service

import (
	"context"

	"github.com/azrilpramudia/libre-go/internal/domain/author"
	"github.com/azrilpramudia/libre-go/internal/domain/book"
	"github.com/google/uuid"
)

type BookService struct {
	repo       book.Repository
	authorRepo author.Repository
}

type ImportBookInput struct {
	Title         string
	AuthorName    string
	ISBN          string
	CoverURL      string
	OpenLibraryID string
	Stock         int32
}

func NewBookService(repo book.Repository, authorRepo author.Repository) *BookService {
	return &BookService{repo: repo, authorRepo: authorRepo}
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

func (s *BookService) ImportBook(ctx context.Context, in ImportBookInput) (*book.Book, error) {
	b, err := s.repo.Create(ctx, &book.Book{
		Title:         in.Title,
		ISBN:          in.ISBN,
		CoverURL:      in.CoverURL,
		OpenLibraryID: in.OpenLibraryID,
		Stock:         in.Stock,
	})
	if err != nil {
		return nil, err
	}

	if in.AuthorName != "" {
		a, err := s.authorRepo.FindOrCreateByName(ctx, in.AuthorName)
		if err != nil {
			return nil, err
		}
		if err := s.repo.LinkAuthor(ctx, b.ID, a.ID); err != nil {
			return nil, err
		}
	}

	return b, nil
}
