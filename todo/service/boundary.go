package service

import (
	"context"

	"github.com/cjcjcj/todo/todo/entities"
)

// TodoService represent Todo service interface
type TodoService interface {
	Create(ctx context.Context, item *entities.Todo) error
	Close(ctx context.Context, item *entities.Todo) error
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entities.Todo, error)
	GetAll(ctx context.Context) ([]*entities.Todo, error)
}
