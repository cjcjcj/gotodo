package http

import (
	"github.com/cjcjcj/todo/todo/domains"
	"github.com/cjcjcj/todo/todo/gateways/entities"
	"github.com/cjcjcj/todo/todo/service"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
)

// Create is a handler for todo list item creation
func (h *todoHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	todoRequestItems := new(entities.Todo)
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

	todoItem := domains.NewTodoFromString(todoRequestItems.Title)

	switch err := h.TodoService.Create(ctx, todoItem); err {
	case service.ErrInternal:
		h.logger.Error(
			"Create error",
			zap.Error(err),
		)

		responseTodoStatusInternalServerErrorCounter.Inc()
		return c.JSON(http.StatusInternalServerError, err.Error())
	case service.ErrTodoNotFound:
		h.logger.Debug(
			"Create not found",
			zap.Error(err),
		)

		responseTodoStatusNotFoundCounter.Inc()
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusCreated, entities.TodoFromDomainTodo(todoItem))
}
