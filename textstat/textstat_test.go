package textstat

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTextStat(t *testing.T) {
	tests := map[string]struct{
		s string
		want map[string]int
	}{
		"returns empty stat for empty string": {
			s: "",
			want: map[string]int{},
		},
		"returns empty stat for no words string": {
			s: "  ? !    (),",
			want: map[string]int{},
		},
		"words of any case are counted as same in stat": {
			s: "  word100 ??? wOrD100 ✅ Go WORD100 go Word100 ?? /// gO ✅",
			want: map[string]int{
				"word100": 4,
				"go": 3,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := WordCount(tt.s)

			assert.Equal(t, got, tt.want)
		})
	}
}