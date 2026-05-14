package normalize

import "testing"

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

func TestClean(t *testing.T) {
	type test struct {
		got	 string
		want string
	}

	tests := map[string]test{
		"keeps empty string": {
			got: "",
			want: "",
		},
		"keeps normalized string as is": {
			got: "i love golang",
			want: "i love golang",
		},
		"trims external spaces and tabs; squashes internal spaces and tabs": {
			got: "	word0		word1 word2    word3    ",
			want: "word0 word1 word2 word3",
		},
		"lowercases the string": {
			got: "AaaaA AaaAaAaaAAA aAAaAAa A",
			want: "aaaaa aaaaaaaaaaa aaaaaaa a",
		},
		"keeps mixed symbols in string": {
			got: "Why so serious? 🤡",
			want: "why so serious? 🤡",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := Clean(tt.got)

			assertEqual(t, result, tt.want)
		})
	}
}