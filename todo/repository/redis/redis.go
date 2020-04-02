package redis

import (
	"go.uber.org/zap"

	goredis "github.com/go-redis/redis/v7"
)

// NewRedisTodoRepository returns new redis TodoRepository instance
func NewRedisTodoRepository(
	client *goredis.Client,
	logger *zap.Logger,
) *TodoRepository {
	return &TodoRepository{
		logger: logger,

		client: client,
	}
}

type TodoRepository struct {
	logger *zap.Logger

	client *goredis.Client
}
