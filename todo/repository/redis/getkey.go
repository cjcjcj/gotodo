package redis

import "strings"

func (r *todoRepository) getKey(parts ...string) string {
	return strings.Join(parts, separator)
}
