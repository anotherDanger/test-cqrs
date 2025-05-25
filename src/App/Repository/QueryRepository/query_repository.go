package queryrepository

import (
	"context"
	domain "test-cqrs/src/Domain"
)

type QueryRepository interface {
	GetBook(ctx context.Context, params string) ([]*domain.Domain, error)
	GetBookByTitle(ctx context.Context, title string) ([]*domain.Domain, error)
}
