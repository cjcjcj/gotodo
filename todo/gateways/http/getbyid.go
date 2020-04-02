package http

import (
	"github.com/cjcjcj/todo/todo/gateways/entities"
	"github.com/cjcjcj/todo/todo/service"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
)

// GetByID is a handler for receiving todo list item by ID
func (h *todoHandler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	todoDomain, err := h.TodoService.GetByID(ctx, id)
	switch err {
	case service.ErrInternal:
		h.logger.Error(
			"GetById error",
			zap.Error(err),
		)

		responseTodoStatusInternalServerErrorCounter.Inc()
		return c.JSON(http.StatusInternalServerError, err.Error())
	case service.ErrTodoNotFound:
		h.logger.Debug(
			"GetById not found error",
			zap.Error(err),
		)

		responseTodoStatusNotFoundCounter.Inc()
		return c.JSON(http.StatusNotFound, err.Error())
	}

	responseTodoStatusOKCounter.Inc()
	return c.JSON(http.StatusOK, entities.TodoFromDomainTodo(todoDomain))
}
