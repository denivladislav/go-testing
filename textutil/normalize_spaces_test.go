package textutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeSpaces(t *testing.T) {
	tests := map[string]struct{
		s string
		want string
	}{
		"normalizes empty string to empty string": {
			s: "",
			want: "",
		},
		"normalizes spaces and tabs string to empty string": {
			s: "   		",
			want: "",
		},
		"trims outer spaces; normalizes mixed string": {
			s: "	😎 cool guy ?   погнали ",
			want: "😎 cool guy ? погнали",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := NormalizeSpaces(tt.s)

			assert.Equal(t, tt.want, got)
		})
	}
}