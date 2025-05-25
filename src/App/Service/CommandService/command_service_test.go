package commandservice

import (
	"context"
	"errors"
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

func TestCommandAddBookFailedAllEmpty(t *testing.T) {
	_, db, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	db.ExpectBegin()
	svc := mocks.NewCommandService(t)

	svc.On("AddBook", mock.Anything, mock.Anything).Return(nil, errors.New("body cannot empty"))
	db.ExpectRollback()
	_, err = svc.AddBook(context.Background(), nil)
	if err == nil {
		t.Error(err)
	}

	assert.Equal(t, "body cannot empty", err.Error())
}

func TestCommandBookFailedIdEmpty(t *testing.T) {
	_, db, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	input := &domain.Domain{
		Author: "Tester 1",
		Title:  "Test 1",
		Genre:  "Testing 1",
	}

	db.ExpectBegin()

	svc := mocks.NewCommandService(t)
	svc.On("AddBook", mock.Anything, mock.Anything).Return(nil, errors.New("id cannot empty"))

	_, err = svc.AddBook(context.Background(), input)

	db.ExpectRollback()

	assert.Equal(t, "id cannot empty", err.Error())
}

func TestCommandBookFailedAuthorEmpty(t *testing.T) {
	_, db, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	input := &domain.Domain{
		Title: "Test 1",
		Genre: "Testing 1",
	}

	db.ExpectBegin()
	svc := mocks.NewCommandService(t)
	svc.On("AddBook", mock.Anything, mock.Anything).Return(nil, errors.New("author cannot empty"))
	_, err = svc.AddBook(context.Background(), input)
	db.ExpectRollback()

	assert.Equal(t, errors.New("author cannot empty"), err)
}

func TestCommandBookFailedTitleEmpty(t *testing.T) {
	_, db, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	input := &domain.Domain{
		Author: "Tester 1",
		Genre:  "Testing",
	}

	db.ExpectBegin()
	svc := mocks.NewCommandService(t)
	svc.On("AddBook", mock.Anything, mock.Anything).Return(nil, errors.New("title cannot empty"))
	_, err = svc.AddBook(context.Background(), input)
	db.ExpectRollback()

	assert.Equal(t, errors.New("title cannot empty"), err)
}

func TestCommandFailedGenreEmpty(t *testing.T) {
	_, db, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	input := &domain.Domain{
		Author: "Tester 1",
		Title:  "Test 1",
	}

	db.ExpectBegin()
	svc := mocks.NewCommandService(t)
	svc.On("AddBook", mock.Anything, mock.Anything).Return(nil, errors.New("genre cannot empty"))
	_, err = svc.AddBook(context.Background(), input)
	db.ExpectRollback()

	assert.Equal(t, errors.New("genre cannot empty"), err)
}
