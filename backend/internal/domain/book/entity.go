package book

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID uuid.UUID
	Title	string
	ISBN	string
	OpenLibraryID	string
	CoverURL	string
	CategoryID	*uuid.UUID
	Stock	int32
	Available	int32
	CreatedAt	time.Time
	UpdatedAt	time.Time
}