package queryservice

import (
	"context"
	helpers "test-cqrs/src/App/Helpers"
	queryrepository "test-cqrs/src/App/Repository/QueryRepository"
	domain "test-cqrs/src/Domain"

	"github.com/sirupsen/logrus"
)

type QueryServiceImpl struct {
	repo queryrepository.QueryRepository
}

func NewQueryServiceImpl(repo queryrepository.QueryRepository) QueryService {
	return &QueryServiceImpl{
		repo: repo,
	}
}

func (svc *QueryServiceImpl) GetBook(ctx context.Context, key string, params string) ([]*domain.Domain, error) {
	result, err := svc.repo.GetBook(ctx, key, params)
	if err != nil {
		helpers.NewErr("src/App/logs/service", logrus.ErrorLevel, err)
		return nil, err
	}

	return result, nil
}
