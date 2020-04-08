package todo

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
	"github.com/cjcjcj/todo/todo/service/errors"
)

func (s *todoService) Create(
	ctx context.Context,
	item *domains.Todo,
) (*domains.Todo, error) {
	todo, err := s.repo.Create(ctx, item)
	if err != nil {
		return nil, errors.ErrInternal
	}

	return todo, nil
}
