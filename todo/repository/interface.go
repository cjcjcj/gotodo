package repository

import (
	"context"
	"github.com/cjcjcj/todo/todo/repository/entities"
)

// TodoRepository represent the Todo's repository contract
type TodoRepository interface {
	Create(ctx context.Context, item *entities.Todo) (*entities.Todo, error)
	Update(ctx context.Context, item *entities.Todo) (*entities.Todo, error)
	Delete(ctx context.Context, ID string) error

	GetByID(ctx context.Context, ID string) (*entities.Todo, error)
}