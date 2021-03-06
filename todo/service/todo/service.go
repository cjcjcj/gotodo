package todo

import "github.com/cjcjcj/todo/todo/repository"

type todoService struct {
	repo repository.TodoRepository
}

// NewTodoService is a constructor for todo service
func NewTodoService(todoRepo repository.TodoRepository) *todoService {
	return &todoService{
		repo: todoRepo,
	}
}
