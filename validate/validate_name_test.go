package validate

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

func validateError(t *testing.T, err error, checkErr func(err error) bool) {
	t.Helper()

	if !checkErr(err) {
		t.Errorf("error check failed for error: %v", err)
	}
}

func TestValidateNameEmptyString(t *testing.T) {
	checkErr := func(err error) bool {
		return errors.Is(err, ErrEmptyName)
	}

	err := ValidateName("")
	validateError(t, err, checkErr)
}

func TestValidateNameNotEmptyString(t *testing.T) {
	err := ValidateName("some string")

	assertError(t, err, false)
}