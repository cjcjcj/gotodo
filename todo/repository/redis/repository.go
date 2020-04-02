package redis

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type todoRepository struct {
	logger *zap.Logger

	Conn redis.Conn
}
