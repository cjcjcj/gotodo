package redis

import (
	"context"
	"go.uber.org/zap"
)

func (r *TodoRepository) Delete(ctx context.Context, ID string) error {
	r.logger.Debug(
		"deleting TODO item",
		zap.String("id", ID),
	)

	var (
		key = r.getKey(redisTODOField, ID)
	)

	cmd := r.client.Del(key)
	if _, err := cmd.Result(); err != nil {
		r.logger.Error(
			"TODO item delete operation error",
			zap.String("id", ID),
			zap.Error(err),
		)

		return err
	}

	return nil
}
