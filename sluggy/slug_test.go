package sluggy

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSlug(t *testing.T) {
	tests := map[string]struct{
		s string
		want string
	}{
		"empty string is processed as empty": {
			s: "",
			want: "",
		},
		"outer spaces trimmed; inner spaces are squashed to gyphen": {
			s: "  hello   darkness      ",
			want: "hello-darkness",
		},
		"several gyphens are squashed to one and diacritics removed": {
			s: "my----old- , ?!  friend",
			want: "my-old-friend",
		},
		"string is lowecased": {
			s: "AbcDEfgHIj",
			want: "abcdefghij",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Slug(tt.s)

			assert.Equal(t, tt.want, got)
		})
	}
}