package redis

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
	"github.com/cjcjcj/todo/todo/repository/errors"
	"go.uber.org/zap"
)

func (r *todoRepository) Update(
	ctx context.Context,
	todo *domains.Todo,
) error {
	r.logger.Debug(
		"updating TODO item",
		zap.String("id", todo.ID),
	)

	var (
		key = r.getKey(redisTODOField, todo.ID)
	)

	cmd := r.client.Set(key, todo, 0)
	if _, err := cmd.Result(); err != nil {
		r.logger.Error(
			"TODO item update operation error",
			zap.Any("source_item", todo),
			zap.Any("to_update_item", todo),
			zap.Error(err),
		)

		return errors.InternalError
	}

	return nil
}
