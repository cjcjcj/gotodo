package service

import (
	"context"
	"errors"
	"github.com/cjcjcj/todo/todo/repository/entities"

	"github.com/cjcjcj/todo/todo/domains"
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

func (s *todoService) Create(ctx context.Context, item *domains.Todo) error {
	itemE := entities.NewTodo(item.ID, item.Title, item.Closed)
	err := s.TodoRepo.Create(ctx, itemE)

	if err != nil {
		return ErrInternal
	}
	return nil
}

func (s *todoService) Close(ctx context.Context, item *domains.Todo) error {
	item.Closed = true
	itemE := entities.NewTodo(item.ID, item.Title, item.Closed)
	err := s.TodoRepo.Update(ctx, itemE)

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

func (s *todoService) GetByID(ctx context.Context, id uint) (*domains.Todo, error) {
	ve, err := s.TodoRepo.GetByID(ctx, id)
	v := domains.NewTodo(ve.ID, ve.Title, ve.Closed)

	if err != nil {
		return nil, ErrInternal
	}
	if v == nil {
		return nil, ErrTodoNotFound
	}

	return v, nil
}

func (s *todoService) GetAll(ctx context.Context) ([]*domains.Todo, error) {
	ves, err := s.TodoRepo.GetAll(ctx)
	vs := make([]*domains.Todo, 0)
	for _, ve := range ves {
		vs = append(vs, domains.NewTodo(ve.ID, ve.Title, ve.Closed))
	}

	if err != nil {
		return nil, ErrInternal
	}
	if len(vs) == 0 {
		return vs, ErrTodoNotFound
	}

	return vs, nil
}
