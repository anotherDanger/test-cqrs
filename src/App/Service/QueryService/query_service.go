package queryservice

import (
	"context"
	domain "test-cqrs/src/Domain"
)

type QueryService interface {
	GetBook(ctx context.Context, params string) ([]*domain.Domain, error)
}
