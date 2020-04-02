package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initEcho() *echo.Echo {
	e := echo.New()

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return e
}
