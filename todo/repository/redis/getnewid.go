package redis

import (
	"context"
	"go.uber.org/zap"
	"strconv"
)

func (r *todoRepository) getNewID(ctx context.Context) (string, error) {
	var (
		key = r.getKey(redisIDCounterField)
	)

	cmd := r.client.Incr(key)
	id, err := cmd.Result()
	if err != nil {
		r.logger.Error(
			"ID increment failed",
			zap.Error(err),
		)
		return "", err
	}

	return strconv.Itoa(int(id)), nil
}
