package codeeval

import (
	"fmt"
	"sort"
)

type CharSeq []rune

// These functions are for sort.Sort()
func (a CharSeq) Len() int           { return len(a) }
func (a CharSeq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CharSeq) Less(i, j int) bool { return a[i] < a[j] }

/* WARNING: NOT CORRECT!!!!!!!!!!!!!
 */
func BurrowsWheelerDecomp(in string) string {
	if len(in) < 1 {
		return ""
	}
	comparr := string2arr(in)
	sortedarr := String2ArrSorted(in)

	idx := findnext(sortedarr, '$')
	// sortedarr = append(sortedarr[idx:], sortedarr[:idx]...)
	// comparr = append(comparr[idx:], comparr[:idx]...)

	resultarr := make([]rune, len(comparr))
	resultarr[idx] = sortedarr[idx] // $
	c := comparr[idx]
	sortedarr[idx] = 0x0
	for i := 1; i < len(sortedarr)-1; i++ {
		pos := findnext(sortedarr, c)
		if pos < 0 {
			return string(resultarr)
		}
		resultarr[pos] = c
		c = comparr[pos]
		sortedarr[pos] = 0x0
	}
	resultarr[0] = c

	// xmap := make(map[rune]rune)
	// for pos, val := range sortedarr {
	// 	xmap[val] = comparr[pos]
	// }
	// decomparr := make([]rune, len(comparr))
	// for i := 0; i < len(sortedarr); i++ {
	// 	val := sortedarr[i]
	// 	decomparr[i] = xmap[val]
	// 	delete(xmap, val)
	// }

	return string(resultarr)
}

func String2ArrSorted(in string) []rune {
	arr := string2arr(in)
	sort.Sort(CharSeq(arr))
	// idx := findnext(arr, '$')
	// arr = append(arr[idx:], arr[:idx]...)
	return arr
}

func string2arr(in string) []rune {
	arr := make([]rune, len(in))
	for pos, char := range in {
		arr[pos] = rune(char)
	}
	return arr
}

func arr2string(arr []rune) string {
	return fmt.Sprintf("%v", arr)
}

func findnext(arr []rune, c rune) int {
	for pos, val := range arr {
		if val == c {
			return pos
		}
	}
	return -1
}
