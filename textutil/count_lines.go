package textutil

import (
	"bufio"
	"io"
)

// CountLines считает количество строк в ридере (по разделителю \n).
func CountLines(r io.Reader) (int, error) {
	var n int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n++
	}
	return n, scanner.Err()
}