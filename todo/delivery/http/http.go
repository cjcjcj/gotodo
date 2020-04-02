package http

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/cjcjcj/todo/todo/entities"
	"github.com/cjcjcj/todo/todo/service"
)

// Todo represents todo entity which comes w/ http requests
type Todo struct {
	ID     uint   `json:"id" query:"id"`
	Title  string `json:"title" validate:"required"`
	Closed bool   `json:"closed"`
}

// TodoFromDomainTodo is a copy constructor for Todo
func TodoFromDomainTodo(td *entities.Todo) *Todo {
	return &Todo{
		ID:     td.ID,
		Title:  td.Title,
		Closed: td.Closed,
	}
}

type todoHandler struct {
	logger *zap.Logger

	TodoService service.TodoService
}

func newTodoHandler(
	todoService service.TodoService,
	logger *zap.Logger,
) *todoHandler {
	return &todoHandler{
		logger:      logger,
		TodoService: todoService,
	}
}

// InitializeTodoHTTPHandler starts HTTP Handler for todo delivery
func InitializeTodoHTTPHandler(
	e *echo.Echo,
	s service.TodoService,
	logger *zap.Logger,
) {
	handler := newTodoHandler(s, logger)

	e.GET("/todo", handler.GetAll)
	e.POST("/todo", handler.Create)
	e.GET("/todo/:id", handler.GetByID)
	e.DELETE("/todo/:id", handler.Delete)
	e.PUT("/todo/:id", handler.Close)
}

// GetAll is a handler for receiving all todo list items
func (h *todoHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	domainTodos, err := h.TodoService.GetAll(ctx)
	switch err {
	case service.ErrInternal:
		h.logger.Error(
			"GetAll error",
			zap.Error(err),
		)

		responseTodoStatusInternalServerErrorCounter.Inc()
		return c.JSON(http.StatusInternalServerError, err.Error())
	case service.ErrTodoNotFound:
		h.logger.Debug(
			"GetAll not found",
			zap.Error(err),
		)

		responseTodoStatusNotFoundCounter.Inc()
		return c.JSON(http.StatusNotFound, err.Error())
	}

	var result []*Todo
	for _, domainTodo := range domainTodos {
		result = append(result, TodoFromDomainTodo(domainTodo))
	}

	responseTodoStatusOKCounter.Inc()
	return c.JSON(http.StatusOK, result)
}

// Create is a handler for todo list item creation
func (h *todoHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	todoRequestItems := new(Todo)
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

	todoItem := entities.NewTodo(todoRequestItems.Title)

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

	return c.JSON(http.StatusCreated, TodoFromDomainTodo(todoItem))
}

// GetByID is a handler for receiving todo list item by ID
func (h *todoHandler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(
			"GetById error",
			zap.Error(err),
		)

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoDomain, err := h.TodoService.GetByID(ctx, uint(id))
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
	return c.JSON(http.StatusOK, TodoFromDomainTodo(todoDomain))
}

// Delete is a handler for removing todo list item
func (h *todoHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(
			"Delete error",
			zap.Error(err),
		)

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	switch err = h.TodoService.Delete(ctx, uint(id)); err {
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

// Close is a handler for marking todo list item as done
func (h *todoHandler) Close(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(
			"Close error",
			zap.Error(err),
		)

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoDomain, err := h.TodoService.GetByID(ctx, uint(id))
	switch err {
	case service.ErrInternal:
		h.logger.Error(
			"Close error",
			zap.Error(err),
		)

		responseTodoStatusInternalServerErrorCounter.Inc()
		return c.JSON(http.StatusInternalServerError, err.Error())
	case service.ErrTodoNotFound:
		h.logger.Debug(
			"Close not found",
			zap.Error(err),
		)

		responseTodoStatusNotFoundCounter.Inc()
		return c.JSON(http.StatusNotFound, err.Error())
	}

	switch err = h.TodoService.Close(ctx, todoDomain); err {
	case service.ErrInternal:
		h.logger.Error(
			"Close error",
			zap.Error(err),
		)

		responseTodoStatusInternalServerErrorCounter.Inc()
		return c.JSON(http.StatusInternalServerError, err.Error())
	case service.ErrTodoNotFound:
		h.logger.Debug(
			"Close not found error",
			zap.Error(err),
		)

		responseTodoStatusNotFoundCounter.Inc()
		return c.JSON(http.StatusNotFound, err.Error())
	}

	responseTodoStatusOKCounter.Inc()
	return c.JSON(http.StatusOK, TodoFromDomainTodo(todoDomain))
}
