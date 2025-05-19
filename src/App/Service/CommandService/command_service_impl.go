package commandservice

import (
	"context"
	"database/sql"
	helpers "test-cqrs/src/App/Helpers"
	commandrepository "test-cqrs/src/App/Repository/CommandRepository"
	domain "test-cqrs/src/Domain"

	"github.com/sirupsen/logrus"
)

type CommandServiceImpl struct {
	tx   *sql.DB
	repo commandrepository.CommandRepository
}

func NewCommandServiceImpl(tx *sql.DB, repo commandrepository.CommandRepository) CommandService {
	return &CommandServiceImpl{}
}

func (svc *CommandServiceImpl) AddBook(ctx context.Context, entity *domain.Domain) (*domain.Domain, error) {
	tx, err := svc.tx.Begin()
	if err != nil {
		helpers.NewErr("../commandservice", logrus.ErrorLevel, err)
		return nil, err
	}

	result, err := svc.repo.AddBook(ctx, tx, entity)
	if err != nil {
		helpers.NewErr("../commandservice", logrus.ErrorLevel, err)
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return result, nil

}
