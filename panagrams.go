package codeeval

import (
	_ "fmt"
	"strings"
)

func Panagram(in string) string {
	in = strings.ToLower(in)
	present := make(map[byte]bool)
	for i := 0; i < len(in); i++ {
		present[in[i]] = true
	}
	result := make([]byte, 0)
	letter := 'a'

	for i := 0; i < 26; i++ {
		if !present[byte(letter)] {
			result = append(result, byte(letter))
		}
		letter = letter + 1
	}

	// 	for k, _ := range present {
	// 		result = append(result, k)
	// 	}
	// fmt.Println(result)
	return string(result)
}
