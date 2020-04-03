package redis

import (
	"go.uber.org/zap"

	goredis "github.com/go-redis/redis/v7"
)

type todoRepository struct {
	logger *zap.Logger

	client *goredis.Client
}

// NewRedisTodoRepository returns new redis todoRepository instance
func NewRedisTodoRepository(
	client *goredis.Client,
	logger *zap.Logger,
) *todoRepository {
	return &todoRepository{
		logger: logger,

		client: client,
	}
}
