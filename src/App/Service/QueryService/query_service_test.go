package queryservice

import (
	"context"
	"errors"
	"test-cqrs/src/App/Service/QueryService/mocks"
	domain "test-cqrs/src/Domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQueryServiceFindAll(t *testing.T) {
	svc := mocks.NewQueryService(t)
	id := uuid.New()
	book := []*domain.Domain{
		{
			Id:     id,
			Author: "Tester 1",
			Title:  "Test 1",
			Genre:  "Test",
		},
	}

	svc.On("GetBook", mock.Anything, mock.Anything, mock.Anything).Return(book, nil)
	result, _ := svc.GetBook(context.Background(), "Author", "Test 1")

	assert.Equal(t, "Test 1", result[0].Title)
}

func TestQueryServiceNull(t *testing.T) {
	svc := mocks.NewQueryService(t)

	svc.On("GetBook", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("cannot find book"))
	_, err := svc.GetBook(context.Background(), "title", "testttttttttttttt")
	assert.Equal(t, errors.New("cannot find book"), err)
}
