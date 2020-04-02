package entities

// Todo represents todo in redis data model
type Todo struct {
	ID     string
	Title  string
	Closed bool
}

func NewTodoFromTodo(t *Todo) *Todo {
	return &Todo{
		ID:     t.ID,
		Title:  t.Title,
		Closed: t.Closed,
	}
}
