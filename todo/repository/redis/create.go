package redis

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
	"go.uber.org/zap"
)

func (r *todoRepository) Create(
	ctx context.Context,
	todo *domains.Todo,
) (*domains.Todo, error) {
	r.logger.Debug(
		"creating new TODO item",
		zap.Any("item", todo),
	)

	ID, err := r.getNewID(ctx)
	if err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("item", todo),
			zap.Error(err),
		)

		return nil, err
	}

	var (
		key = r.getKey(redisTODOField, ID)
	)

	todo.ID = ID
	cmd := r.client.Set(key, todo, 0)
	if _, err := cmd.Result(); err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("source_item", todo),
			zap.Any("to_insert_item", todo),
			zap.Error(err),
		)

		return nil, err
	}

	return todo, nil
}
