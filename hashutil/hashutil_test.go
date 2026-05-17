package hashutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSlug(t *testing.T) {
	tests := map[string]struct{
		s string
		want string
	}{
		"empty string is processed correctly": {
			s: "",
			want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		"simple string is processed correctly": {
			s: "bonjour",
			want: "2cb4b1431b84ec15d35ed83bb927e27e8967d75f4bcd9cc4b25c8d879ae23e18",
		},
		"unicode is processed correctly": {
			s: "приветики",
			want: "df2c04a38d4a26ed0560806805eb9f8211eb4cb86888667c48b960272e907d08",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := HashSHA256(tt.s)

			assert.Equal(t, got, tt.want)
		})
	}
}