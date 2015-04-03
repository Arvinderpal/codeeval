package codeeval

import (
	"testing"
)

func TestPanagram(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"abcdefghijklmnopqrstuvwxy", "z"},
		{"abcdefghijklmnopqrstuvwxyz", ""},
		{"xyz", "abcdefghijklmnopqrstuvw"},
		{"A quick brown fox jumps over the lazy dog", ""},
		{"A slow yellow fox crawls under the proactive dog", "bjkmqz"},
	}
	for _, c := range cases {
		got := Panagram(c.in)
		if got != c.want {
			t.Errorf("Panagram(%v) == %v, want %v ",
				c.in, got, c.want)
		}
	}
}
