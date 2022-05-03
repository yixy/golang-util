package str

import (
	"strings"
)

// HasPrefixes returns true if the string s has any of the given prefixes.
func HasPrefixes(src string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(src, prefix) {
			return true
		}
	}
	return false
}

func Match(src string, mathers ...string) bool {
	for _, math := range mathers {
		if src == math {
			return true
		}
	}
	return false
}
