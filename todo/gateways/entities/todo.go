package entities

import "github.com/cjcjcj/todo/todo/domains"

// Todo represents todo entity which comes w/ http requests
type Todo struct {
	ID     string `json:"id" query:"id"`
	Title  string `json:"title" validate:"required"`
	Closed bool   `json:"closed"`
}

func (t *Todo) ToDomainTodo() *domains.Todo {
	return domains.NewTodo(t.ID, t.Title, t.Closed)
}

// NewTodoFromDomainTodo is a copy constructor for Todo
func NewTodoFromDomainTodo(td *domains.Todo) *Todo {
	return &Todo{
		ID:     td.ID,
		Title:  td.Title,
		Closed: td.Closed,
	}
}
