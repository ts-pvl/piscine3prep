package piscine

func CountAlpha(s string) int {
	count := 0
	for _, r := range s {
		// если бука то ++
		// if r >= '0' && r <= '9' {}
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count = count + 1 // count++
		}
	}
	return count
}
