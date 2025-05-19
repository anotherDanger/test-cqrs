package commandrepository

import (
	"context"
	"database/sql"
	domain "test-cqrs/src/Domain"
)

type CommandRepository interface {
	AddBook(ctx context.Context, tx *sql.Tx, entity *domain.Domain) (*domain.Domain, error)
}
