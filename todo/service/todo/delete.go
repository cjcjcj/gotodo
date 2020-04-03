package todo

import (
	"context"
	"github.com/cjcjcj/todo/todo/service/errors"
)

func (s *todoService) Delete(ctx context.Context, ID string) error {
	err := s.TodoRepo.Delete(ctx, ID)

	if err != nil {
		return errors.ErrInternal
	}
	return nil
}
