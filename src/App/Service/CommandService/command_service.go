package commandservice

import (
	"context"
	domain "test-cqrs/src/Domain"
)

type CommandService interface {
	AddBook(ctx context.Context, entity *domain.Domain) (*domain.Domain, error)
}
