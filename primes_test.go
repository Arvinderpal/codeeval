package codeeval

import (
	"testing"
)

func TestGetPrimes(t *testing.T) {
	cases := []struct {
		in   int
		want []int
	}{
		{1, nil},
		{-1, nil},
		{2, []int{2}},
		{3, []int{2, 3}},
		{4, []int{2, 3}},
		{5, []int{2, 3, 5}},
		{10, []int{2, 3, 5, 7}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		// {1000, []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}
	for _, c := range cases {
		got := GetPrimes(c.in)
		eq := func(got, want []int) bool {
			if len(got) != len(want) {
				return false
			}
			for i := 0; i < len(got); i++ {
				if got[i] != want[i] {
					return false
				}
			}
			return true
		}
		if !eq(got, c.want) {
			t.Errorf("GetPrimes(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCheckprime_naive(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{13, true},
		{9, false},
		{929, true},
		{98689, true},
		{-1, false},
		{9989899, true},
		// {9999987899999, true},
	}
	for _, c := range cases {
		got := Checkprime_naive(c.in)
		if got != c.want {
			t.Errorf("Checkprime_naive(%v) == %q, want %q", c.in, got, c.want)
		}
	}
}
