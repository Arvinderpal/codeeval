package codeeval

func GetPrimes(n int) []int {
	if n < 2 {
		return nil
	}
	primes := make([]int, 1)
	primes[0] = 2
	for i := 3; i <= n; i++ {
		if Checkprime_naive(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func Checkprime_naive(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num&0x1 == 0 {
		return false
	}
	for i := 3; i < num/2; i = i + 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
