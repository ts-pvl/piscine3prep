package piscine

func FindPrevPrime(nb int) int {
	for i := nb; i >= 2; i-- { // 5 -> 2
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
