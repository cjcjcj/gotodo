package entities

// Todo represents todo in redis data model
type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Closed bool   `json:"closed"`
}

func NewTodo(id uint, title string, closed bool) *Todo{
	return &Todo{
		ID:     id,
		Title:  title,
		Closed: closed,
	}
}