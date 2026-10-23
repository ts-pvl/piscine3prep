package piscine

func DigitLen(n int, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}

	count := 0 // сколько раз поделить
	for n > 0 {
		n = n / base
		count++ // count = count + 1
	}
	return count
}
