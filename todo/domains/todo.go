package domains

import (
	"github.com/cjcjcj/todo/todo/repository/entities"
)

// Todo defines todo model
type Todo struct {
	ID     string
	Title  string
	Closed bool
}

func (t *Todo) ToRepoTodo() *entities.Todo {
	return &entities.Todo{
		ID:     t.ID,
		Title:  t.Title,
		Closed: t.Closed,
	}
}

func NewTodo(ID string, title string, closed bool) *Todo {
	return &Todo{
		ID:     ID,
		Title:  title,
		Closed: closed,
	}
}

// NewTodoFromString returns new Todo instance
func NewTodoFromString(title string) *Todo {
	return NewTodo("", title, false)
}

func NewTodoFromRepoTodo(todo *entities.Todo) *Todo {
	return &Todo{
		ID:     todo.ID,
		Title:  todo.Title,
		Closed: todo.Closed,
	}
}
