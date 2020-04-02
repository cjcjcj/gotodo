package service

import (
	"context"
	"errors"
	repoErrors "github.com/cjcjcj/todo/todo/repository/errors"

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
func NewTodoService(todoRepo repository.TodoRepository) *todoService {
	return &todoService{TodoRepo: todoRepo}
}

func (s *todoService) Create(ctx context.Context, item *domains.Todo) error {
	repoItem := item.ToRepoTodo()

	if _, err := s.TodoRepo.Create(ctx, repoItem); err != nil {
		return ErrInternal
	}
	return nil
}

func (s *todoService) Close(ctx context.Context, item *domains.Todo) error {
	item.Closed = true

	repoItem := item.ToRepoTodo()
	if _, err := s.TodoRepo.Update(ctx, repoItem); err != nil {
		return ErrInternal
	}
	return nil
}

func (s *todoService) Delete(ctx context.Context, ID string) error {
	err := s.TodoRepo.Delete(ctx, ID)

	if err != nil {
		return ErrInternal
	}
	return nil
}

func (s *todoService) GetByID(ctx context.Context, ID string) (*domains.Todo, error) {
	repoTodo, err := s.TodoRepo.GetByID(ctx, ID)

	switch err.(type) {
	case nil:
		domainTodo := domains.NewTodoFromRepoTodo(repoTodo)
		return domainTodo, nil
	case *repoErrors.NotFoundError:
		return nil, ErrTodoNotFound
	default:
		// any other errors
		return nil, ErrInternal
	}
}
