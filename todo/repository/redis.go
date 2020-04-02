package repository

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"

	"github.com/cjcjcj/todo/todo/entities"
	"github.com/gomodule/redigo/redis"
)

const (
	redisTODOsField = "hm"
	redisIDfield    = "id"
)

// Todo represents todo in redis data model
type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Closed bool   `json:"closed"`
}

func newTodoFromEntityTodo(te *entities.Todo) *Todo {
	return &Todo{
		ID:     te.ID,
		Title:  te.Title,
		Closed: te.Closed,
	}
}

func (tr *Todo) toEntityTodo() *entities.Todo {
	return &entities.Todo{
		ID:     tr.ID,
		Title:  tr.Title,
		Closed: tr.Closed,
	}
}

type redisTodoRepository struct {
	logger *zap.Logger

	Conn redis.Conn
}

// NewRedisTodoRepository returns new redis todo repository instance
func NewRedisTodoRepository(
	conn redis.Conn,
	logger *zap.Logger,
) TodoRepository {
	return &redisTodoRepository{
		logger: logger,

		Conn: conn,
	}
}

func (r *redisTodoRepository) Create(ctx context.Context, te *entities.Todo) error {
	r.logger.Debug(
		"creating new TODO item",
		zap.Any("item", te),
	)

	tr := newTodoFromEntityTodo(te)

	newID, err := redis.Int(r.Conn.Do("INCR", redisIDfield))
	if err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("item", te),
			zap.Error(err),
		)

		return err
	}

	tr.ID = uint(newID)

	trB, err := json.Marshal(tr)
	if err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("item", te),
			zap.Error(err),
		)

		return err
	}

	if _, err = r.Conn.Do("HSET", redisTODOsField, tr.ID, string(trB)); err != nil {
		r.logger.Error(
			"TODO create item operation error",
			zap.Any("item", te),
			zap.Error(err),
		)

		return err
	}

	*te = *tr.toEntityTodo()

	return nil
}

func (r *redisTodoRepository) Update(ctx context.Context, te *entities.Todo) error {
	r.logger.Debug(
		"updating TODO item",
		zap.Uint("id", te.ID),
	)
	tr := newTodoFromEntityTodo(te)

	trB, err := json.Marshal(tr)
	if err != nil {
		r.logger.Error(
			"TODO item update operation error",
			zap.Uint("id", tr.ID),
			zap.Error(err),
		)

		return err
	}

	if _, err = r.Conn.Do("HSET", redisTODOsField, tr.ID, string(trB)); err != nil {
		r.logger.Error(
			"TODO item update operation error",
			zap.Uint("id", tr.ID),
			zap.Error(err),
		)

		return err
	}

	*te = *tr.toEntityTodo()

	return nil
}

func (r *redisTodoRepository) Delete(ctx context.Context, id uint) error {
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

func (r *redisTodoRepository) GetByID(ctx context.Context, id uint) (*entities.Todo, error) {
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

	tr := new(Todo)
	err = json.Unmarshal(vb, tr)
	if err != nil {
		return nil, err
	}

	return tr.toEntityTodo(), nil
}

func (r *redisTodoRepository) GetAll(ctx context.Context) ([]*entities.Todo, error) {
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
	tr := new(Todo)

	for _, trB := range v {
		err = json.Unmarshal(trB, tr)
		if err != nil {
			r.logger.Error(
				"receiving all todo items error",
				zap.Error(err),
			)

			return nil, err
		}
		todos = append(todos, tr.toEntityTodo())
	}

	return todos, nil
}
