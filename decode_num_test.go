package codeeval

import (
	"testing"
)

func TestDecodeNum(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1", 1},
		{"12", 2},
		{"123", 3},
		{"1213", 5},
		{"12113", 8},
		{"43", 1},
		{"423", 2},
		{"999", 1},
		{"987654", 1},
		// {"121212", 3},
	}
	for _, c := range cases {
		got := DecodeNum(c.in)
		if got != c.want {
			t.Errorf("Checkpalindrom(%v) == %v, want %v ",
				c.in, got, c.want)
		}
	}
}
