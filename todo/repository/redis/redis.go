package redis

import (
	"context"
	"github.com/cjcjcj/todo/todo/repository/entities"
	"github.com/cjcjcj/todo/todo/repository/errors"
	"go.uber.org/zap"
	"strconv"

	goredis "github.com/go-redis/redis/v7"
)

func (r *TodoRepository) getNewID(ctx context.Context) (string, error) {
	var (
		key = r.getKey(redisIDCounterField)
	)

	cmd := r.client.Incr(key)
	id, err := cmd.Result()
	if err != nil {
		r.logger.Error(
			"ID increment failed",
			zap.Error(err),
		)
		return "", err
	}

	return strconv.Itoa(int(id)), nil
}

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

func (r *TodoRepository) Delete(ctx context.Context, ID string) error {
	r.logger.Debug(
		"deleting TODO item",
		zap.String("id", ID),
	)

	var (
		key = r.getKey(redisTODOField, ID)
	)

	cmd := r.client.Del(key)
	if _, err := cmd.Result(); err != nil {
		r.logger.Error(
			"TODO item delete operation error",
			zap.String("id", ID),
			zap.Error(err),
		)

		return err
	}

	return nil
}

func (r *TodoRepository) GetByID(
	ctx context.Context,
	ID string,
) (*entities.Todo, error) {
	r.logger.Debug(
		"receiving TODO item",
		zap.String("id", ID),
	)

	var (
		key = r.getKey(redisTODOField, ID)

		result = &entities.Todo{}
	)

	switch err := r.client.Get(key).Scan(result); err {
	case nil:
		// OK
		r.logger.Debug(
			"TODO item successfully received",
			zap.String("id", ID),
		)

		break
	case goredis.Nil:
		// if value not present
		return nil, &errors.NotFoundError{ID: ID}
	default:
		// any other error
		r.logger.Error(
			"getting TODO item error",
			zap.String("id", ID),
			zap.Error(err),
		)

		return nil, err
	}

	return result, nil
}
