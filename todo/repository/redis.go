package repository

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"

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
	Conn redis.Conn
}

// NewRedisTodoRepository returns new redis todo repository instance
func NewRedisTodoRepository(conn redis.Conn) TodoRepository {
	return &redisTodoRepository{Conn: conn}
}

func (r *redisTodoRepository) Create(ctx context.Context, te *entities.Todo) error {
	logrus.Info("adding todo item w/")
	tr := newTodoFromEntityTodo(te)

	newID, err := redis.Int(r.Conn.Do("INCR", redisIDfield))
	if err != nil {
		logrus.WithField("err", err).Warn("adding todo item w/")
		return err
	}

	tr.ID = uint(newID)

	trB, err := json.Marshal(tr)
	if err != nil {
		logrus.WithFields(logrus.Fields{"id": tr.ID, "err": err}).Warn("adding todo item w/")
		return err
	}

	if _, err = r.Conn.Do("HSET", redisTODOsField, tr.ID, string(trB)); err != nil {
		logrus.WithFields(logrus.Fields{"id": tr.ID, "err": err}).Warn("adding todo item w/")
		return err
	}

	*te = *tr.toEntityTodo()

	return nil
}

func (r *redisTodoRepository) Update(ctx context.Context, te *entities.Todo) error {
	logrus.WithField("id", te.ID).Info("updating todo item w/")
	tr := newTodoFromEntityTodo(te)

	trB, err := json.Marshal(tr)
	if err != nil {
		logrus.WithFields(logrus.Fields{"id": te.ID, "err": err}).Warn("updating todo item w/")
		return err
	}

	if _, err = r.Conn.Do("HSET", redisTODOsField, tr.ID, string(trB)); err != nil {
		logrus.WithFields(logrus.Fields{"id": tr.ID, "err": err}).Warn("updating todo item w/")
		return err
	}

	*te = *tr.toEntityTodo()

	return nil
}

func (r *redisTodoRepository) Delete(ctx context.Context, id uint) error {
	logrus.WithField("id", id).Info("removing todo item w/")
	if _, err := r.Conn.Do("HDEL", redisTODOsField, id); err != nil {
		logrus.WithFields(logrus.Fields{"id": id, "err": err}).Warn("on removing todo item error occurred")
		return err
	}

	return nil
}

func (r *redisTodoRepository) GetByID(ctx context.Context, id uint) (*entities.Todo, error) {
	logrus.WithField("id", id).Info("getting todo item w/")

	vb, err := redis.Bytes(r.Conn.Do("HGET", redisTODOsField, id))
	switch err {
	case nil:
		// OK
		logrus.WithField("id", id).Info("getted todo item w/")
		break
	case redis.ErrNil:
		// if value not present
		return nil, nil
	default:
		// any other error
		logrus.WithFields(logrus.Fields{"id": id, "err": err}).Warn("getting todo item w/")
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
	logrus.Info("getting all todo items")
	v, err := redis.ByteSlices(r.Conn.Do("HVALS", redisTODOsField))
	if err != nil {
		logrus.WithField("err", err).Warn("getting all todo items w/")
		return nil, err
	}

	var todos []*entities.Todo
	tr := new(Todo)

	for _, trB := range v {
		err = json.Unmarshal(trB, tr)
		if err != nil {
			logrus.WithField("err", err).Info("getting all todo items w/")
			return nil, err
		}
		todos = append(todos, tr.toEntityTodo())
	}

	return todos, nil
}
