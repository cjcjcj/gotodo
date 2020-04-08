package todo

import (
	"context"
	mock_repository "github.com/cjcjcj/todo/mocks/github.com/cjcjcj/todo/todo/repository"
	repoErrors "github.com/cjcjcj/todo/todo/repository/errors"
	serviceErrors "github.com/cjcjcj/todo/todo/service/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_todoService_Delete(t *testing.T) {
	const (
		id = "1"
	)

	ctx := context.TODO()

	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		mockRepo.EXPECT().Delete(gomock.Any(), gomock.Eq(id)).Return(nil)

		err := s.Delete(ctx, id)
		assert.NoError(t, err)
	})

	t.Run("ERR", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockTodoRepository(ctrl)
		s := NewTodoService(mockRepo)

		mockRepo.EXPECT().Delete(gomock.Any(), gomock.Eq(id)).Return(repoErrors.InternalError)

		err := s.Delete(ctx, id)
		assert.EqualError(t, err, serviceErrors.ErrInternal.Error())
	})

}
