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

func Test_todoService_Close(t *testing.T) {
	const (
		id    = "1"
		title = "test"
	)

	ctx := context.TODO()

	t.Run("OK-opened-to-closed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		gotFromRepo := &domains.Todo{id, title, false}
		sentToRepo := &domains.Todo{id, title, true}

		mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(id)).Return(gotFromRepo, nil)
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Eq(sentToRepo)).Return(nil)

		got, err := s.Close(ctx, id)
		assert.NoError(t, err)

		expected := &domains.Todo{id, title, true}
		assert.Equal(t, *got, *expected, "should be equal")
	})

	t.Run("OK-closed-to-closed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		gotFromRepo := &domains.Todo{id, title, true}
		sentToRepo := &domains.Todo{id, title, true}

		mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(id)).Return(gotFromRepo, nil)
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Eq(sentToRepo)).Return(nil)

		got, err := s.Close(ctx, id)
		assert.NoError(t, err)

		expected := &domains.Todo{id, title, true}
		assert.Equal(t, *got, *expected, "should be equal")
	})

	t.Run("err-get-by-id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		errRepo := &repoErrors.NotFoundError{}
		mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(id)).Return(nil, errRepo)
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Times(0)

		got, err := s.Close(ctx, id)
		assert.EqualError(t, err, serviceErrors.ErrTodoNotFound.Error())
		assert.Nil(t, got)
	})

	t.Run("err-update", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		gotFromRepo := &domains.Todo{id, title, true}
		errRepo := repoErrors.InternalError
		mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Eq(id)).Return(gotFromRepo, nil)
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errRepo)

		got, err := s.Close(ctx, id)
		assert.EqualError(t, err, serviceErrors.ErrInternal.Error())
		assert.Nil(t, got)
	})
}
