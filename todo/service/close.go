package service

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
	"github.com/cjcjcj/todo/todo/service/errors"
)

func (s *todoService) Close(ctx context.Context, item *domains.Todo) error {
	item.Closed = true

	repoItem := item.ToRepoTodo()
	if _, err := s.TodoRepo.Update(ctx, repoItem); err != nil {
		return errors.ErrInternal
	}
	return nil
}
