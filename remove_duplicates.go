package codeeval

import (
	"sort"
)

// RemoveDuplicatesAndSort will remove all duplicate entries in the input slice and return a new slice that is sorted
func RemoveDuplicatesAndSort(s []int) []int {
	var newSlice []int
	dupSet := make(map[int]bool)
	for _, e := range s {
		if dupSet[e] != true {
			dupSet[e] = true
		}
	}
	for k, _ := range dupSet {
		newSlice = append(newSlice, k)
	}
	sort.Ints(newSlice)
	return newSlice
}
