package main

import (
	"fmt"
	delivery "github.com/cjcjcj/todo/todo/delivery/http"
	"github.com/cjcjcj/todo/todo/repository"
	"github.com/cjcjcj/todo/todo/service"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"io"
	"os"
	"sort"
)

var (
	commit    = "unknown"
	version   = "unknown"
	buildDate = "unknown"
)

const (
	logsDir       = "/var/log/todo"
	logrusPath    = logsDir + "/logs.log"
	accesslogPath = logsDir + "/access.log"
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

	logrus.Info(redisAddr)
	logrus.Info(echoAddr)

	lrfd, err := os.OpenFile(logrusPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return cli.NewMultiError(err)
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

	redisConn, err := redis.Dial(
		"tcp",
		redisAddr,
		redis.DialDatabase(cfg.Redis.DB),
	)
	if err != nil {
		logrus.Fatalf("redis; %v, %s", err, redisAddr)
	}
	defer redisConn.Close()

	e := initEcho(echoLogFd)

	todoRepo := repository.NewRedisTodoRepository(redisConn)
	todoService := service.NewTodoService(todoRepo)
	delivery.InitializeTodoHTTPHandler(e, todoService)

	logrus.Error(e.Start(echoAddr))

	return
}
