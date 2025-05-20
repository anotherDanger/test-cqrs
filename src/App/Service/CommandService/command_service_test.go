package commandservice

import (
	"context"
	"test-cqrs/src/App/Service/CommandService/mocks"
	domain "test-cqrs/src/Domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCommandAddBookSuccess(t *testing.T) {
	_, sql, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	id := uuid.New()

	input := &domain.Domain{
		Id:     id,
		Author: "Tester 1",
		Title:  "Test 1",
		Genre:  "Testing",
	}

	expected := &domain.Domain{
		Id:     id,
		Author: "Tester 1",
		Title:  "Test 1",
		Genre:  "Testing",
	}

	svc := mocks.NewCommandService(t)
	sql.ExpectBegin()
	svc.On("AddBook", mock.Anything, mock.Anything).Return(expected, nil)
	sql.ExpectCommit()

	result, err := svc.AddBook(context.Background(), input)
	if err != nil {
		t.Error(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
