package codeeval

import (
	"testing"
)

func TestReverseK(t *testing.T) {
	cases := []struct {
		in   string
		in2  int
		want string
	}{
		{"12345", 2, "21435"},
		{"12345", 3, "32145"},
		{"12345", 5, "54321"},
		{"12345239575730-274356456456", 5, "5432175932-0375534724654656"},
	}
	for _, c := range cases {
		got := ReverseK(c.in, c.in2)
		if got != c.want {
			t.Errorf("ReverseK(%v, %v) == %v, want %v \n",
				c.in, c.in2, got, c.want)
		}
	}
}

func TestCheckRotation(t *testing.T) {
	cases := []struct {
		in   string
		in2  string
		want bool
	}{
		{"Hello", "Hello", true},
		{"Hello", "lloHe", true},
		{"Basefont", "tBasefon", true},
		{"Basefont", "tBasfeon", false},
		{"Basefont", "xBasfeon", false},
		{"nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnna", "annnnnnnnnnnnnnnnnnnnnnnnnnnnnnn", true},
		{"nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnna", "nnnnnnnnnnnnnnannnnnnnnnnnnnnnnn", true},
		{"nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba", "annnnnnnnnnnnnnnnnnnnnnnnnnnnnnnbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", true},
		{"", "", true},
	}
	for _, c := range cases {
		got := CheckRotation(c.in, c.in2)
		if got != c.want {
			t.Errorf("CheckRotation(%v, %v) == %v, want %v \n",
				c.in, c.in2, got, c.want)
		}
	}
}
