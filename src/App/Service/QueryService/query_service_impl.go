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

func (svc *QueryServiceImpl) GetBook(ctx context.Context, params string) ([]*domain.Domain, error) {
	result, err := svc.repo.GetBook(ctx, params)
	if err != nil {
		helpers.NewErr("/home/andhikadanger/cqrs/src/App/logs/queryservice", logrus.ErrorLevel, err)
		return nil, err
	}

	return result, nil
}
