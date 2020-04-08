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

func Test_todoService_Create(t *testing.T) {
	const (
		id    = "1"
		title = "test"
	)

	ctx := context.TODO()

	t.Run("OK-closed-false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		const closed = false

		toCreate := &domains.Todo{"", title, closed}
		created := &domains.Todo{id, title, closed}

		mockRepo.EXPECT().Create(gomock.Any(), gomock.Eq(toCreate)).Return(created, nil)

		got, err := s.Create(ctx, toCreate)
		assert.NoError(t, err)

		expected := &domains.Todo{id, title, closed}
		assert.Equal(t, *got, *expected, "should be equal")
	})

	t.Run("OK-closed-true", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		const closed = true

		toCreate := &domains.Todo{"", title, closed}
		created := &domains.Todo{id, title, closed}

		mockRepo.EXPECT().Create(gomock.Any(), gomock.Eq(toCreate)).Return(created, nil)

		got, err := s.Create(ctx, toCreate)
		assert.NoError(t, err)

		expected := &domains.Todo{id, title, closed}
		assert.Equal(t, *got, *expected, "should be equal")
	})

	t.Run("ERR-create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		const closed = true

		toCreate := &domains.Todo{"", title, closed}

		mockRepo.EXPECT().Create(gomock.Any(), gomock.Eq(toCreate)).Return(nil, repoErrors.InternalError)

		got, err := s.Create(ctx, toCreate)
		assert.EqualError(t, err, serviceErrors.ErrInternal.Error())
		assert.Nil(t, got)
	})
}
