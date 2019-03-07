package codeeval

import (
	"fmt"
	"testing"
)

func TestMReceiversOneSender(t *testing.T) {

	cases := []struct {
		M int
	}{
		{1},
		{5},
		{10},
	}

	for _, c := range cases {
		MReceiversOneSender(c.M)
		fmt.Printf("\n\n")
	}
}

func TestMReceiversNSenders(t *testing.T) {

	cases := []struct {
		M int
		N int
	}{
		{5, 10},
	}
	for _, c := range cases {
		MReceiversNSenders(c.M, c.N)
	}
}
