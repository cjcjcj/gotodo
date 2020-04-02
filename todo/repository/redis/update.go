package redis

import (
	"context"
	"github.com/cjcjcj/todo/todo/repository/entities"
	"go.uber.org/zap"
)

func (r *TodoRepository) Update(
	ctx context.Context,
	todo *entities.Todo,
) (*entities.Todo, error) {
	r.logger.Debug(
		"updating TODO item",
		zap.String("id", todo.ID),
	)

	var (
		key = r.getKey(redisTODOField, todo.ID)

		toupdate = entities.NewTodoFromTodo(todo)
	)

	cmd := r.client.Set(key, toupdate, 0)
	if _, err := cmd.Result(); err != nil {
		r.logger.Error(
			"TODO item update operation error",
			zap.Any("source_item", todo),
			zap.Any("to_update_item", toupdate),
			zap.Error(err),
		)

		return nil, err
	}

	return toupdate, nil
}
