package codeeval

import (
	// "fmt"
	"strings"
)

func ReverseK(in string, n int) string {
	if len(in) < 2 {
		return in
	}
	if n < 1 {
		return in
	}

	trune := []rune(in)
	for i := 0; i+n <= len(trune); i += n / 2 {
		for j := i + n - 1; j >= i; i, j = i+1, j-1 {
			trune[i], trune[j] = trune[j], trune[i]
		}
	}
	return string(trune)
}

func CheckRotation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	var (
		xor1 byte
		xor2 byte
	)
	for i := 0; i < len(str1); i++ {
		xor1 ^= byte(str1[i])
		xor2 ^= byte(str2[i])
	}
	// fmt.Printf("xor: %v, %v", xor1, xor2)
	if xor1 != xor2 {

		return false
	}

	str2 += str2
	if strings.Contains(str2, str1) {
		return true
	}
	return false
}
