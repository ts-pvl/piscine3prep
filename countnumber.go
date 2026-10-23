package piscine

func CountNumber(s string) int {
	count := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			count = count + 1
		}
	}
	return count
}
