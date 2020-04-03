package http

import (
	"github.com/cjcjcj/todo/todo/service"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type todoHandler struct {
	logger *zap.Logger

	todoService service.TodoService
}

func newTodoHandler(
	todoService service.TodoService,
	logger *zap.Logger,
) *todoHandler {
	return &todoHandler{
		logger: logger,

		todoService: todoService,
	}
}

// InitializeTodoHandler starts HTTP Handler for todo delivery
func InitializeTodoHandler(
	e *echo.Echo,
	s service.TodoService,
	logger *zap.Logger,
) {
	handler := newTodoHandler(s, logger)

	e.POST("/todo", handler.Create)
	e.GET("/todo/:id", handler.GetByID)
	e.DELETE("/todo/:id", handler.Delete)
	e.PUT("/todo/:id", handler.Close)
}
