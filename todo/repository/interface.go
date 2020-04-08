package repository

import (
	"context"
	"github.com/cjcjcj/todo/todo/domains"
)

// TodoRepository represent the Todo's repository contract
type TodoRepository interface {
	Create(ctx context.Context, item *domains.Todo) (*domains.Todo, error)
	Update(ctx context.Context, item *domains.Todo) error
	Delete(ctx context.Context, ID string) error

	GetByID(ctx context.Context, ID string) (*domains.Todo, error)
}
