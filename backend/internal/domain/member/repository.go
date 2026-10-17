package member

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, m *Member) (*Member, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Member, error)
	GetByEmail(ctx context.Context, email string) (*Member, error)
	List(ctx context.Context, limit, offset int32) ([]*Member, error)
}
