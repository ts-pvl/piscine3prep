package piscine

func FirstWord(s string) string {
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			break
		} else {
			word = word + string(s[i])
		}
	}
	return word
}
