package errors

import "errors"

var (
	ErrTodoNotFound = errors.New("ToDo item not found")
	ErrInternal     = errors.New("internal error")
)
