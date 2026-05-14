package even

import "testing"

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

func TestIsEven(t *testing.T) {
	type test struct {
		number	 int
		expected bool
	}

	tests := map[string]test{
		"determines positive even number": {
			number: 1002,
			expected: true,
		},
		"determines negative odd number": {
			number: -11,
			expected: false,
		},
		"treats zero as even number": {
			number: 0,
			expected: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := IsEven(tt.number)

			assertEqual(t, result, tt.expected)
		})
	}
}
