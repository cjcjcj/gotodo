package domains

// Todo defines todo model
type Todo struct {
	ID     string
	Title  string
	Closed bool
}

func (t *Todo) Close() {
	t.Closed = true
}

func (t *Todo) Open() {
	t.Closed = false
}

func NewTodo(ID string, title string, closed bool) *Todo {
	return &Todo{
		ID:     ID,
		Title:  title,
		Closed: closed,
	}
}
