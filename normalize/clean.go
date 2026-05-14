package normalize

import "strings"

// Clean trims spaces, collapses internal spaces to one
// and lowercases ASCII letters.
func Clean(s string) string {
	tokens := strings.Fields(s)
	if len(tokens) == 0 {
		return ""
	}
	return strings.ToLower(strings.Join(tokens, " "))
}
