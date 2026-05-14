package safe

import "testing"

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

func validateRecover(t *testing.T, recover any, checkRecover func(r any) bool) {
	t.Helper()

	if recover == nil {
		t.Fatalf("recover expected")
	}

	if !checkRecover(recover) {
		t.Errorf("recover check failed for recover: %v", recover)
	}
}

func TestMustAtNegativeIndex(t *testing.T) {
	checkRecover := func(recover any) bool {
		return recover == PanicOutOfRange
	}

	defer func() {
		r := recover()

		validateRecover(t, r, checkRecover)
	}()

	MustAt([]int{0, 1}, -1)
}

func TestMustAtExceedingIndex(t *testing.T) {
	checkRecover := func(recover any) bool {
		return recover == PanicOutOfRange
	}

	defer func() {
		r := recover()

		validateRecover(t, r, checkRecover)
	}()

	MustAt([]int{0, 1}, 2)
}

func TestMustAt(t *testing.T) {
	s := []int{10, 20}
	res1 := MustAt(s, 0)
	res2 := MustAt(s, 1)

	assertEqual(t, res1, 10)
	assertEqual(t, res2, 20)
}