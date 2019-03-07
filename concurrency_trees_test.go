package codeeval

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestTreesEqual(t *testing.T) {
	cases := []struct {
		t1  *tree.Tree
		t2  *tree.Tree
		exp bool
	}{
		{tree.New(2), tree.New(2), true},
		{tree.New(2), tree.New(3), false},
	}

	for _, c := range cases {
		out := TreesEqual(c.t1, c.t2)
		if out != c.exp {
			t.Errorf("expected %s to equal %s but got %v", c.t1, c.t2, out)
		}
	}
}
