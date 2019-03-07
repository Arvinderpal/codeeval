package codeeval

import "testing"

func TestRemoveAllOccurances(t *testing.T) {
	cases := []struct {
		s   []int
		o   []int
		exp []int
	}{
		{[]int{2, 3, 4, 5, 6}, []int{2, 6}, []int{3, 4, 5}},
		{[]int{2}, []int{2}, []int{}},
		{[]int{}, []int{2}, []int{}},
	}
	for _, c := range cases {
		n := RemoveAllOccurances(c.s, c.o)
		if !Equal(c.exp, n) {
			t.Errorf("expected %v but got %v \n", c.exp, n)
		}
	}
}
