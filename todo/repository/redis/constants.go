package redis

const (
	rootKey   = "todo"
	separator = ":"

	redisTODOField      = rootKey + separator + "todo"
	redisIDCounterField = rootKey + separator + "ID"
)
