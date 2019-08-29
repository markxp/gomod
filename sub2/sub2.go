package sub2

import (
	"strings"
)

// Sub is an alias to strings.ReplaceAll
func Sub(i, rep, sub string) string {
	return strings.ReplaceAll(i, rep, sub)
}
