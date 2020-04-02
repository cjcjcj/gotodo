package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
)

type echoValidator struct {
	validator *validator.Validate
}

func (ev *echoValidator) Validate(i interface{}) error {
	return ev.validator.Struct(i)
}

func initEcho(logOutput io.Writer) *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: logOutput}))
	e.Validator = &echoValidator{validator: validator.New()}

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return e
}
