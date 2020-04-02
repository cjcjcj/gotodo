package main

import (
	"github.com/urfave/cli"
)

var cliFlags = []cli.Flag{
	// REDIS
	cli.StringFlag{
		Name:   "todo-redis-host",
		Usage:  "Redis host",
		EnvVar: "TODO_REDIS_HOST",
	},
	cli.StringFlag{
		Name:   "todo-redis-port",
		Usage:  "Redis port",
		EnvVar: "TODO_REDIS_PORT",
	},
	cli.IntFlag{
		Name:   "todo-redis-db",
		Usage:  "Redis db",
		EnvVar: "TODO_REDIS_DB",
	},
}

func newConfigurationFromContext(
	ctx *cli.Context,
) (
	config *configuration,
	err error,
) {
	config = &configuration{
		Redis: configRedis{
			Host: ctx.String("todo-redis-host"),
			Port: ctx.String("todo-redis-port"),
			DB:   ctx.Int("todo-redis-db"),
		},
	}

	return
}

// configuration represents application configuration store.
type configuration struct {
	Redis configRedis `valid:"required"`
}

type configRedis struct {
	Host string `valid:"required"`
	Port string `valid:"required"`
	DB   int    `valid:"required"`
}
