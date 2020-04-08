package todo

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
	repoErrors "github.com/cjcjcj/todo/todo/repository/errors"
	"github.com/cjcjcj/todo/todo/service/errors"
)

func (s *todoService) GetByID(
	ctx context.Context,
	ID string,
) (*domains.Todo, error) {
	todo, err := s.repo.GetByID(ctx, ID)

	switch err.(type) {
	case nil:
		return todo, nil
	case *repoErrors.NotFoundError:
		return nil, errors.ErrTodoNotFound
	default:
		// any other errors
		return nil, errors.ErrInternal
	}
}
