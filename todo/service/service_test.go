package service

import (
	"context"
	"testing"

	"github.com/cjcjcj/todo/mocks"
	"github.com/cjcjcj/todo/todo/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTodoCreate(t *testing.T) {
	const (
		id     = uint(0)
		title  = "test"
		closed = false
	)

	mockRepo := &mocks.TodoRepository{}

	te := entities.NewTodo(title)
	expectedTe := entities.Todo{
		ID:     id,
		Title:  title,
		Closed: closed,
	}

	t.Run("OK", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, te).Once().Return(nil)

		s := NewTodoService(mockRepo)

		err := s.Create(context.TODO(), te)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)

		assert.Equal(t, *te, expectedTe, "should be equal")
	})
}

func TestTodoClose(t *testing.T) {
	const (
		id    = uint(1)
		title = "test"
	)

	mockRepo := &mocks.TodoRepository{}
	te := &entities.Todo{
		ID:     id,
		Title:  title,
		Closed: false,
	}
	expectedTe := entities.Todo{
		ID:     id,
		Title:  title,
		Closed: true,
	}

	t.Run("OK", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything, te).Once().Return(nil)

		s := NewTodoService(mockRepo)

		err := s.Close(context.TODO(), te)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		assert.Equal(t, *te, expectedTe, "should be equal")
	})
}

func TestTodoDelete(t *testing.T) {
	const id = uint(0)

	mockRepo := &mocks.TodoRepository{}

	t.Run("OK", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything, id).Once().Return(nil)

		s := NewTodoService(mockRepo)

		err := s.Delete(context.TODO(), id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestTodoGetById(t *testing.T) {
	const (
		id     = uint(1)
		title  = "test"
		closed = false
	)

	mockRepo := &mocks.TodoRepository{}

	expectedTe := entities.Todo{
		ID:     id,
		Title:  title,
		Closed: closed,
	}

	t.Run("OK", func(t *testing.T) {
		mockRepo.On("GetByID", mock.Anything, id).Once().Return(&expectedTe, nil)

		s := NewTodoService(mockRepo)

		te, err := s.GetByID(context.TODO(), id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		assert.Equal(t, *te, expectedTe, "should be equal")
	})
}
