package service

import (
	"context"

	"github.com/cjcjcj/todo/todo/domains"
)

// TodoService represent Todo service interface
type TodoService interface {
	Create(ctx context.Context, item *domains.Todo) (*domains.Todo, error)
	Close(ctx context.Context, ID string) (*domains.Todo, error)
	Delete(ctx context.Context, ID string) error

	GetByID(ctx context.Context, ID string) (*domains.Todo, error)
}
