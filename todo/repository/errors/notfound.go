package errors

import "fmt"

type NotFoundError struct {
	ID string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found error: %v", e.ID)
}
