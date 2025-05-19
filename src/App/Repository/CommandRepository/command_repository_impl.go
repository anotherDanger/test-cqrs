package commandrepository

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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
	entity.Id = id
	_, err := tx.ExecContext(ctx, query, id, entity.Author, entity.Title, entity.Genre)
	if err != nil {
		helpers.NewErr("../logs/commandrepository", logrus.ErrorLevel, err)
		return nil, err
	}

	reqBody, err := json.Marshal(entity)
	if err != nil {
		helpers.NewErr("../logs/commandrepository", logrus.FatalLevel, err)
		return nil, err
	}

	_, err = http.Post(fmt.Sprintf("http://localhost:9200/books/_create/%s", id), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		helpers.NewErr("../logs/commandrepository", logrus.FatalLevel, err)
		return nil, err
	}

	return entity, nil

}
