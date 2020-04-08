package todo

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
	"github.com/cjcjcj/todo/todo/service/errors"
)

func (s *todoService) Close(
	ctx context.Context,
	ID string,
) (*domains.Todo, error) {
	item, err := s.GetByID(ctx, ID)
	if err != nil {
		return nil, errors.ErrTodoNotFound
	}
	item.Close()

	if err := s.repo.Update(ctx, item); err != nil {
		return nil, errors.ErrInternal
	}
	return item, nil
}
