package repository

import (
	"context"
	"github.com/cjcjcj/todo/todo/repository/entities"
)

// TodoRepository represent the Todo's repository contract
type TodoRepository interface {
	Create(ctx context.Context, item *entities.Todo) error
	Update(ctx context.Context, item *entities.Todo) error
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entities.Todo, error)
	GetAll(ctx context.Context) ([]*entities.Todo, error)
}
