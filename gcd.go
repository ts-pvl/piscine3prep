package piscine

func Gcd(a uint, b uint) uint {
	var min uint = 0
	if a < b {
		min = a
	} else {
		min = b
	}
	for i := min; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}
