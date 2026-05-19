package textutil

import (
	"strings"
)

// NormalizeSpaces схлопывает последовательности пробельных символов в один пробел
// и обрезает крайние пробелы.
func NormalizeSpaces(s string) string {
	tokens := strings.Fields(s)
	return strings.Join(tokens, " ")
}