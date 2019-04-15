package main

import (
	"io"
	"os"

	"github.com/go-playground/validator"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	delivery "github.com/cjcjcj/todo/todo/delivery/http"
	"github.com/cjcjcj/todo/todo/repository"
	"github.com/cjcjcj/todo/todo/service"
)

const (
	logsDir       = "/var/log/todo"
	logrusPath    = logsDir + "/logs.log"
	accesslogPath = logsDir + "/access.log"
)

const (
	defaultEchoAddr  = ":8080"
	defaultRedisAddr = "127.0.0.1:6379"
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

func main() {
	var (
		redisAddr = os.Getenv("TODO_REDIS_ADDR")
		echoAddr  = os.Getenv("TODO_ECHO_ADDR")
	)
	if redisAddr == "" {
		// redis on localhost w/ default port
		redisAddr = defaultRedisAddr
	}
	if echoAddr == "" {
		echoAddr = defaultEchoAddr
	}

	logrus.Info(redisAddr)
	logrus.Info(echoAddr)

	lrfd, err := os.OpenFile(logrusPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatal(err)
	}
	defer lrfd.Close()
	logrus.SetOutput(lrfd)

	var echoLogFd io.Writer
	fd, err := os.Create(accesslogPath)
	if err != nil {
		logrus.Warn(err)
		echoLogFd = os.Stdout
	} else {
		echoLogFd = fd
		defer fd.Close()
	}

	redisConn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		logrus.Fatalf("redis; %v, %s", err, redisAddr)
	}
	defer redisConn.Close()

	e := initEcho(echoLogFd)

	todoRepo := repository.NewRedisTodoRepository(redisConn)
	todoService := service.NewTodoService(todoRepo)
	delivery.InitializeTodoHTTPHandler(e, todoService)

	logrus.Error(e.Start(echoAddr))
}
