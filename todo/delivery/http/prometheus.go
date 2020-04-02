package http

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	responseTodoStatusOKCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_todo_response_status_ok",
		Help: "response todo status OK counter"})

	responseTodoStatusNotFoundCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_todo_response_status_not_found",
		Help: "response todo status not found counter"})

	responseTodoStatusInternalServerErrorCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_todo_response_status_internal_server_error",
		Help: "response todo status internal server errror counter"})
)
