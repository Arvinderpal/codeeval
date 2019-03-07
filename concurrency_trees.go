package codeeval

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// walk will do a in-order walk through the tree writing elements to the channel along the way
func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
	// fmt.Printf("%v ", t.Value)
}

// Walk will use the walk() func to do the in-order walk of the tree and will close dCh when done
func Walk(t *tree.Tree, ch chan int, dCh chan struct{}) {

	walk(t, ch)
	dCh <- struct{}{}
}

// TreesEqual will compare the in-order list of nodes in the trees and return true if they are equal
func TreesEqual(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	ch2 := make(chan int)
	dCh1 := make(chan struct{})
	dCh2 := make(chan struct{})
	var done1, done2 bool

	go Walk(t1, ch1, dCh1)
	go Walk(t2, ch2, dCh2)

	var s1, s2 []int
LOOP:
	for {
		select {
		case v := <-ch1:
			s1 = append(s1, v)
		case v := <-ch2:
			s2 = append(s2, v)
		case <-dCh1:
			done1 = true
			if done2 == true {
				break LOOP
			}
		case <-dCh2:
			done2 = true
			if done1 == true {
				break LOOP
			}
		}
	}
	fmt.Printf("t1 %v\n", s1)
	fmt.Printf("t2 %v\n", s2)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
