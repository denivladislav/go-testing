package hello

import (
	"errors"
	"testing"
)

func assertError(t *testing.T, err error, wantErr bool) {
	t.Helper()

	if (err != nil) != wantErr {
		t.Errorf("error = %v, wantErr = %v", err, wantErr)
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

func validateError(t *testing.T, err error, checkErr func(err error) bool) {
	t.Helper()

	if !checkErr(err) {
		t.Errorf("error check failed for error: %v", err)
	}
}

func TestHelloEmptyName(t *testing.T) {
	checkErr := func(err error) bool {
		return errors.Is(err, ErrEmptyName)
	}

	_, err := Hello("")
	validateError(t, err, checkErr)
}

func TestHello(t *testing.T) {
	type test struct {
		got string
		want string
	}

	tests := map[string]test{
		"basic string": {
			got: "Go",
			want: "Hello, Go",
		},
		"external spaces are not trimmed": {
			got: "  Go ",
			want: "Hello,   Go ",
		},
		"unicode is processed correctly": {
			got: "Гоферок 😎",
			want: "Hello, Гоферок 😎",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := Hello(tt.got)
			assertError(t, err, false)

			assertEqual(t, result, tt.want)
		})
	}
}