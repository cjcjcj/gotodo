package http

import (
	"github.com/cjcjcj/todo/todo/service"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
)

// Delete is a handler for removing todo list item
func (h *todoHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	switch err := h.TodoService.Delete(ctx, id); err {
	case service.ErrInternal:
		h.logger.Error(
			"Delete error",
			zap.Error(err),
		)

		responseTodoStatusInternalServerErrorCounter.Inc()
		return c.JSON(http.StatusInternalServerError, err.Error())
	case service.ErrTodoNotFound:
		h.logger.Debug(
			"Delete error",
			zap.Error(err),
		)

		responseTodoStatusNotFoundCounter.Inc()
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
