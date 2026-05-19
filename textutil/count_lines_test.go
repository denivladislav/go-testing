package textutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestCountLines(t *testing.T) {
	tests := map[string]struct{
		text string
		want int
	}{
		"counts empty text as zero lines": {
			text: "",
			want: 0,
		},
		"counts empty inner lines as lines": {
			text: "\na\n\nb\n",
			want: 4,
		},
		"counts one text line as one": {
			text: "Hello, it's me",
			want: 1,
		},
		"counts lines in text with newlines correctly": {
			text: "	I was wondering\nif after all these years,\nyou'd like to meet\nTo go over everything  ",
			want: 4,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(tt.text)

			got, err := CountLines(r)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}