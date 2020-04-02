package domains

// Todo defines todo model
type Todo struct {
	ID     uint
	Title  string
	Closed bool
}

func NewTodo(ID uint, title string, closed bool) *Todo {
	return &Todo{
		ID:     ID,
		Title:  title,
		Closed: closed,
	}
}

// NewTodoFromString returns new Todo instance
func NewTodoFromString(title string) *Todo {
	return NewTodo(0, title, false)
}
