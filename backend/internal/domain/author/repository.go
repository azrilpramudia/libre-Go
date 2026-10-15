package author

import "context"

type Repository interface {
	FindOrCreateByName(ctx context.Context, name string) (*Author, error)
}
