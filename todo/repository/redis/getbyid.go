package redis

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
	"github.com/cjcjcj/todo/todo/repository/errors"
	"go.uber.org/zap"

	goredis "github.com/go-redis/redis/v7"
)

func (r *todoRepository) GetByID(
	ctx context.Context,
	ID string,
) (*domains.Todo, error) {
	r.logger.Debug(
		"receiving TODO item",
		zap.String("id", ID),
	)

	var (
		key = r.getKey(redisTODOField, ID)

		result = &domains.Todo{}
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

		return nil, errors.InternalError
	}

	return result, nil
}
