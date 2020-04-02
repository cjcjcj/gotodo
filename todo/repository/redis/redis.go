package redis

import (
	"context"
	"encoding/json"
	"github.com/cjcjcj/todo/todo/repository"
	"github.com/cjcjcj/todo/todo/repository/entities"
	"go.uber.org/zap"

	"github.com/gomodule/redigo/redis"
)

// NewRedisTodoRepository returns new redis todo repository instance
func NewRedisTodoRepository(
	conn redis.Conn,
	logger *zap.Logger,
) repository.TodoRepository {
	return &todoRepository{
		logger: logger,

		Conn: conn,
	}
}

func (r *todoRepository) Create(ctx context.Context, te *entities.Todo) error {
	r.logger.Debug(
		"creating new TODO item",
		zap.Any("item", te),
	)

	newID, err := redis.Int(r.Conn.Do("INCR", redisIDfield))
	if err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("item", te),
			zap.Error(err),
		)

		return err
	}
	te.ID = uint(newID)

	trB, err := json.Marshal(te)
	if err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("item", te),
			zap.Error(err),
		)

		return err
	}

	if _, err = r.Conn.Do("HSET", redisTODOsField, te.ID, string(trB)); err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("item", te),
			zap.Error(err),
		)

		return err
	}

	return nil
}

func (r *todoRepository) Update(ctx context.Context, te *entities.Todo) error {
	r.logger.Debug(
		"updating TODO item",
		zap.Uint("id", te.ID),
	)

	trB, err := json.Marshal(te)
	if err != nil {
		r.logger.Error(
			"TODO item update operation error",
			zap.Uint("id", te.ID),
			zap.Error(err),
		)

		return err
	}

	if _, err = r.Conn.Do("HSET", redisTODOsField, te.ID, string(trB)); err != nil {
		r.logger.Error(
			"TODO item update operation error",
			zap.Uint("id", te.ID),
			zap.Error(err),
		)

		return err
	}

	return nil
}

func (r *todoRepository) Delete(ctx context.Context, id uint) error {
	r.logger.Debug(
		"deleting TODO item",
		zap.Uint("id", id),
	)

	if _, err := r.Conn.Do("HDEL", redisTODOsField, id); err != nil {
		r.logger.Error(
			"TODO item delete operation error",
			zap.Uint("id", id),
			zap.Error(err),
		)

		return err
	}

	return nil
}

func (r *todoRepository) GetByID(ctx context.Context, id uint) (*entities.Todo, error) {
	r.logger.Debug(
		"receiving TODO item",
		zap.Uint("id", id),
	)

	vb, err := redis.Bytes(r.Conn.Do("HGET", redisTODOsField, id))
	switch err {
	case nil:
		// OK
		r.logger.Debug(
			"TODO item successfully received",
			zap.Uint("id", id),
		)

		break
	case redis.ErrNil:
		// if value not present
		return nil, nil
	default:
		// any other error
		r.logger.Error(
			"getting TODO item error",
			zap.Uint("id", id),
			zap.Error(err),
		)

		return nil, err
	}

	tr := &entities.Todo{}
	err = json.Unmarshal(vb, tr)
	if err != nil {
		return nil, err
	}

	return tr, nil
}

func (r *todoRepository) GetAll(ctx context.Context) ([]*entities.Todo, error) {
	r.logger.Debug(
		"getting all TODO items",
	)

	v, err := redis.ByteSlices(r.Conn.Do("HVALS", redisTODOsField))
	if err != nil {
		r.logger.Error(
			"receiving all todo items error",
			zap.Error(err),
		)

		return nil, err
	}

	var todos []*entities.Todo
	tr := &entities.Todo{}

	for _, trB := range v {
		err = json.Unmarshal(trB, tr)
		if err != nil {
			r.logger.Error(
				"receiving all todo items error",
				zap.Error(err),
			)

			return nil, err
		}
		todos = append(todos, tr)
	}

	return todos, nil
}
