package tempfiledemo

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func assertError(t *testing.T, err error, wantErr bool) {
	t.Helper()

	if (err != nil) != wantErr {
		t.Errorf("error = %v, wantErr = %v", err, wantErr)
	}
}

func validateError(t *testing.T, err error, checkErr func(err error) bool) {
	t.Helper()

	if !checkErr(err) {
		t.Errorf("error check failed for error: %v", err)
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

func TestWriteLinesToTemp(t *testing.T) {
	customPrefix := "prefix-"
	customLines := []string{
		"line0",
		"line1",
		"line2",
	}

	type test struct {
		prefix string
		lines []string
	}

	tests := map[string]test{
		"processes empty lines": {
			prefix: customPrefix,
			lines: []string{},
		},
		"processes empty prefix": {
			prefix: "",
			lines: customLines,
		},
		"processes lines and prefix": {
			prefix: customPrefix,
			lines: customLines,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				r := recover()
				if r != nil {
					t.Fatalf("panic not expected")
				}
			}()

			filepath, _ := WriteLinesToTemp(tt.prefix, tt.lines)
			data, err := os.ReadFile(filepath)

			assertError(t, err, false)
			assertEqual(t, string(data), strings.Join(tt.lines, "\n"))

			err = os.Remove(filepath)
			assertError(t, err, false)

			_, err = os.Stat(filepath)
			validateError(t, err, func(err error) bool {
				return errors.Is(err, os.ErrNotExist)
			})
		})
	}
}