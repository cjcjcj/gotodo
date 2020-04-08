package todo

import (
	"context"
	"github.com/cjcjcj/todo/todo/service/errors"
)

func (s *todoService) Delete(ctx context.Context, ID string) error {
	err := s.repo.Delete(ctx, ID)

	if err != nil {
		return errors.ErrInternal
	}
	return nil
}
