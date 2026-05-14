package hello

import "errors"

var ErrEmptyName = errors.New("name cannot be empty")

func Hello(name string) (string, error) {
	if name == "" {
		return "", ErrEmptyName
	}
	return "Hello, " + name, nil
}