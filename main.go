package main

import (
	"fmt"
	delivery "github.com/cjcjcj/todo/todo/gateways/http"
	repository "github.com/cjcjcj/todo/todo/repository/redis"
	"github.com/cjcjcj/todo/todo/service/todo"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"sort"

	goredis "github.com/go-redis/redis/v7"
)

var (
	commit    = "unknown"
	version   = "unknown"
	buildDate = "unknown"
)

// main starts program
func main() {
	app := cli.NewApp()
	app.Name = "aleksandr-ilin-go-demo-todo"
	app.Usage = "Aleksandr Ilin golang demo TODO project"
	app.Description = ""
	app.Version = fmt.Sprintf("%s (commit: %s, build date: %s)", version, commit, buildDate)
	app.Action = action

	app.Flags = cliFlags
	sort.Sort(cli.FlagsByName(app.Flags))

	_ = app.Run(os.Args)
}

func action(ctx *cli.Context) (err error) {
	// parse cfg
	cfg, err := newConfigurationFromContext(ctx)
	if err != nil {
		return cli.NewMultiError(err)
	}

	redisAddr := fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port)
	redisClient := goredis.NewClient(&goredis.Options{
		Addr:     redisAddr,
		Password: "",           // no password set
		DB:       cfg.Redis.DB, // use default DB
	})
	echoAddr := ":8080"

	logger := zap.NewExample()
	defer logger.Sync()

	logger.Debug(
		"redis configs: ",
		zap.Any("cfg", cfg.Redis),
	)
	logger.Debug(
		"echo address: ",
		zap.Any("address", echoAddr),
	)

	e := initEcho()

	todoRepo := repository.NewRedisTodoRepository(redisClient, logger)
	todoService := todo.NewTodoService(todoRepo)
	delivery.InitializeTodoHandler(e, todoService, logger)

	return
}

func initEcho() *echo.Echo {
	e := echo.New()

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return e
}
