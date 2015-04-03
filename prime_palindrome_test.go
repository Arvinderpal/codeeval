package codeeval

import (
	"testing"
)

func TestPrimePalindrom(t *testing.T) {
	cases := []struct {
		in    int
		want1 int
		want2 PalindromeError
	}{
		{1, 0, "WTF"},
		{-2, 0, "WTF"},
		{2, 2, ""},
		{3, 3, ""},
		{20, 11, ""},
		{1000, 929, ""},
		{10000, 929, ""},
		{100000, 98689, ""},
		// {10000000, 98689, ""},
	}
	for _, c := range cases {
		got1, got2 := PrimePalindrom(c.in)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Checkpalindrom(%v) == %v, want1 %v | got2=%v want2=%v",
				c.in, got1, c.want1, got2, c.want2)
		}
	}
}

func TestCheckpalindrom(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"9", true},
		{"121", true},
		{"223322", true},
		{"1231", false},
		{"999999999999999999", true},
		{"799123321997", true},
		{"999999999992999999", false},
	}
	for _, c := range cases {
		got := Checkpalindrom(c.in)
		if got != c.want {
			t.Errorf("Checkpalindrom(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
