package http

import (
	"github.com/cjcjcj/todo/todo/gateways/entities"
	serviceErrrors "github.com/cjcjcj/todo/todo/service/errors"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
)

// Close is a handler for marking todo list item as done
func (h *todoHandler) Close(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	todoDomain, err := h.todoService.Close(ctx, id)
	switch err {
	case serviceErrrors.ErrInternal:
		h.logger.Error(
			"Close error",
			zap.Error(err),
		)

		responseTodoStatusInternalServerErrorCounter.Inc()
		return c.JSON(http.StatusInternalServerError, err.Error())
	case serviceErrrors.ErrTodoNotFound:
		h.logger.Debug(
			"Close not found error",
			zap.Error(err),
		)

		responseTodoStatusNotFoundCounter.Inc()
		return c.JSON(http.StatusNotFound, err.Error())
	}

	responseTodoStatusOKCounter.Inc()
	return c.JSON(http.StatusOK, entities.NewTodoFromDomainTodo(todoDomain))
}
