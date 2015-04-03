package codeeval

import (
	"fmt"
)

type PalindromeError string

func (e *PalindromeError) Error() string {
	return fmt.Sprint(e)
}

func Checkpalindrom(in string) bool {
	if len(in) == 1 {
		return true
	} else if len(in) < 1 {
		return false
	}
	for i := 0; i < len(in)/2; i++ {
		if in[i] != in[len(in)-1-i] {
			return false
		}
	}
	return true
}

func PrimePalindrom(n int) (int, PalindromeError) {
	primes := GetPrimes(n)
	err := PalindromeError("")
	for i := len(primes) - 1; i >= 0; i-- {
		str := fmt.Sprint(primes[i])
		if Checkpalindrom(str) {
			return primes[i], err
		}
	}
	err = PalindromeError("WTF")
	return 0, err
}
