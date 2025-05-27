package queryrepository

import (
	"context"
	domain "test-cqrs/src/Domain"
)

type QueryRepository interface {
	GetBook(ctx context.Context, key string, params string) ([]*domain.Domain, error)
}
