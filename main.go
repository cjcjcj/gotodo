package main

import (
	"fmt"
	delivery "github.com/cjcjcj/todo/todo/delivery/http"
	"github.com/cjcjcj/todo/todo/repository"
	"github.com/cjcjcj/todo/todo/service"
	"github.com/gomodule/redigo/redis"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"sort"
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

	app.Run(os.Args)
}

func action(ctx *cli.Context) (err error) {
	// parse cfg
	cfg, err := newConfigurationFromContext(ctx)

	redisAddr := fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port)
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

	redisConn, err := redis.Dial(
		"tcp",
		redisAddr,
		redis.DialDatabase(cfg.Redis.DB),
	)
	if err != nil {
		logger.Error(
			"redis initialization failed",
			zap.Error(err),
		)
		return cli.NewMultiError(err)
	}
	defer redisConn.Close()

	e := initEcho()

	todoRepo := repository.NewRedisTodoRepository(redisConn, logger)
	todoService := service.NewTodoService(todoRepo)
	delivery.InitializeTodoHTTPHandler(e, todoService, logger)

	return
}
