package piscine

func HashCode(dec string) string {
	hash := ""
	for _, r := range dec {
		c := (r + rune(len(dec))) % 127
		if c < 32 {
			c = c + 33
		}
		hash = hash + string(c)
	}
	return hash
}
