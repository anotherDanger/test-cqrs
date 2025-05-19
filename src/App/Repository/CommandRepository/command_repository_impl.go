package commandrepository

import (
	"context"
	"database/sql"
	helpers "test-cqrs/src/App/Helpers"
	domain "test-cqrs/src/Domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CommandRepositoryImpl struct{}

func NewCommandRepositoryImpl() CommandRepository {
	return &CommandRepositoryImpl{}
}

func (repo *CommandRepositoryImpl) AddBook(ctx context.Context, tx *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	query := "insert into books(id, author, title, genre) values(?, ?, ?, ?)"
	id := uuid.New()
	_, err := tx.ExecContext(ctx, query, id, entity.Author, entity.Title, entity.Genre)
	if err != nil {
		helpers.NewErr("../logs/commandrepository", logrus.ErrorLevel, err)
		return nil, err
	}

	return entity, nil

}
