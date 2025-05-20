package queryrepository

import (
	"context"
	domain "test-cqrs/src/Domain"
)

type QueryRepository interface {
	GetBook(ctx context.Context, params string) ([]*domain.Domain, error)
}
