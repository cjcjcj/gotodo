package todo

import (
	"context"
	mock_repository "github.com/cjcjcj/todo/mocks/github.com/cjcjcj/todo/todo/repository"
	"github.com/cjcjcj/todo/todo/domains"
	repoErrors "github.com/cjcjcj/todo/todo/repository/errors"
	serviceErrors "github.com/cjcjcj/todo/todo/service/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_todoService_GetByID(t *testing.T) {
	const (
		id     = "1"
		title  = "test"
		closed = true
	)

	ctx := context.TODO()

	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		gotFromRepo := &domains.Todo{id, title, closed}
		mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(id)).Return(gotFromRepo, nil)

		got, err := s.GetByID(ctx, id)
		assert.NoError(t, err)

		expected := &domains.Todo{id, title, closed}
		assert.Equal(t, *got, *expected, "should be equal")
	})

	t.Run("ERR-404", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		errRepo := &repoErrors.NotFoundError{ID: id}
		mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(id)).Return(nil, errRepo)

		got, err := s.GetByID(ctx, id)
		assert.EqualError(t, err, serviceErrors.ErrTodoNotFound.Error())
		assert.Nil(t, got)
	})

	t.Run("ERR-internal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		errRepo := repoErrors.InternalError
		mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(id)).Return(nil, errRepo)

		got, err := s.GetByID(ctx, id)
		assert.EqualError(t, err, serviceErrors.ErrInternal.Error())
		assert.Nil(t, got)
	})
}
