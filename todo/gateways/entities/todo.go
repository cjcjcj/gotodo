package entities

import "github.com/cjcjcj/todo/todo/domains"

// Todo represents todo entity which comes w/ http requests
type Todo struct {
	ID     string `json:"id" query:"id"`
	Title  string `json:"title" validate:"required"`
	Closed bool   `json:"closed"`
}

// TodoFromDomainTodo is a copy constructor for Todo
func TodoFromDomainTodo(td *domains.Todo) *Todo {
	return &Todo{
		ID:     td.ID,
		Title:  td.Title,
		Closed: td.Closed,
	}
}

