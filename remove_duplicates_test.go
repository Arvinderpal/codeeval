package codeeval

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	cases := []struct {
		input []int
		expec []int
	}{
		{[]int{2, 3, 2, 3}, []int{2, 3}},
		{[]int{2, 2, 2, 2, 2, 2}, []int{2}},
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, c := range cases {
		out := RemoveDuplicatesAndSort(c.input)

		if !Equal(out, c.expec) {
			t.Errorf("expected %v but got %v \n", c.expec, out)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
