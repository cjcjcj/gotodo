package entities

// Todo defines todo model
type Todo struct {
	ID     uint
	Title  string
	Closed bool
}

// NewTodo returns new Todo instance
func NewTodo(title string) *Todo {
	return &Todo{Title: title, Closed: false}
}
