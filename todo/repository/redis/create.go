package redis

import (
	"context"
	"github.com/cjcjcj/todo/todo/repository/entities"
	"go.uber.org/zap"
)

func (r *TodoRepository) Create(
	ctx context.Context,
	todo *entities.Todo,
) (*entities.Todo, error) {
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

	toinsert := entities.NewTodoFromTodo(todo)
	toinsert.ID = ID

	cmd := r.client.Set(key, toinsert, 0)
	if _, err := cmd.Result(); err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("source_item", todo),
			zap.Any("to_insert_item", toinsert),
			zap.Error(err),
		)

		return nil, err
	}

	return toinsert, nil
}
