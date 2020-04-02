package service

import (
	"context"
	"errors"

	"github.com/cjcjcj/todo/todo/entities"
	"github.com/cjcjcj/todo/todo/repository"
)

var (
	ErrTodoNotFound = errors.New("ToDo item not found")
	ErrInternal     = errors.New("internal error")
)

type todoService struct {
	TodoRepo repository.TodoRepository
}

// NewTodoService is a constructor for todo service
func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{TodoRepo: todoRepo}
}

func (s *todoService) Create(ctx context.Context, item *entities.Todo) error {
	err := s.TodoRepo.Create(ctx, item)

	if err != nil {
		return ErrInternal
	}
	return nil
}

func (s *todoService) Close(ctx context.Context, item *entities.Todo) error {
	item.Closed = true
	err := s.TodoRepo.Update(ctx, item)

	if err != nil {
		return ErrInternal
	}
	return nil
}

func (s *todoService) Delete(ctx context.Context, id uint) error {
	err := s.TodoRepo.Delete(ctx, id)

	if err != nil {
		return ErrInternal
	}
	return nil
}

func (s *todoService) GetByID(ctx context.Context, id uint) (*entities.Todo, error) {
	v, err := s.TodoRepo.GetByID(ctx, id)

	if err != nil {
		return nil, ErrInternal
	}
	if v == nil {
		return nil, ErrTodoNotFound
	}

	return v, nil
}

func (s *todoService) GetAll(ctx context.Context) ([]*entities.Todo, error) {
	v, err := s.TodoRepo.GetAll(ctx)

	if err != nil {
		return nil, ErrInternal
	}
	if len(v) == 0 {
		return v, ErrTodoNotFound
	}

	return v, nil
}
