package service

import (
	"context"

	"github.com/cjcjcj/todo/todo/domains"
)

// TodoService represent Todo service interface
type TodoService interface {
	Create(ctx context.Context, item *domains.Todo) error
	Close(ctx context.Context, item *domains.Todo) error
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*domains.Todo, error)
	GetAll(ctx context.Context) ([]*domains.Todo, error)
}
