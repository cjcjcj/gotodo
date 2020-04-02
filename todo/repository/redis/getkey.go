package redis

import "strings"

func (r *TodoRepository) getKey(parts ...string) string {
	return strings.Join(parts, separator)
}
