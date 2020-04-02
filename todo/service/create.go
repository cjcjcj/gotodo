package service

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
	"github.com/cjcjcj/todo/todo/service/errors"
)

func (s *todoService) Create(ctx context.Context, item *domains.Todo) error {
	repoItem := item.ToRepoTodo()

	if _, err := s.TodoRepo.Create(ctx, repoItem); err != nil {
		return errors.ErrInternal
	}
	return nil
}

