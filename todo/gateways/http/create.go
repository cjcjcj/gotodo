package http

import (
	"github.com/cjcjcj/todo/todo/gateways/entities"
	serviceErrors "github.com/cjcjcj/todo/todo/service/errors"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
)

// Create is a handler for todo list item creation
func (h *todoHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	todoRequestItems := &entities.Todo{}
	if err := c.Bind(todoRequestItems); err != nil {
		h.logger.Error(
			"Create error",
			zap.Error(err),
		)

		return err
	}
	if err := c.Validate(todoRequestItems); err != nil {
		h.logger.Error(
			"Create validation error",
			zap.Error(err),
		)

		return err
	}

	todoItem := todoRequestItems.ToDomainTodo()
	todoCreated, err := h.todoService.Create(ctx, todoItem)
	switch err {
	case serviceErrors.ErrInternal:
		h.logger.Error(
			"Create error",
			zap.Error(err),
		)

		responseTodoStatusInternalServerErrorCounter.Inc()
		return c.JSON(http.StatusInternalServerError, err.Error())
	case serviceErrors.ErrTodoNotFound:
		h.logger.Debug(
			"Create not found",
			zap.Error(err),
		)

		responseTodoStatusNotFoundCounter.Inc()
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusCreated, entities.NewTodoFromDomainTodo(todoCreated))
}
