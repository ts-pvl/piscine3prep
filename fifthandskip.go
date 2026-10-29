package piscine

// "abcdefghijklmnopqrstuwxyz"

func FifthAndSkip(str string) string {
	runes := []rune(str)

	if len(runes) == 0 {
		return "\n"
	}

	if len(runes) < 5 {
		return "Invalid Input\n"
	}

	count := 0
	result := ""

	for _, r := range str {
		if r == ' ' {
			continue
		}

		count++
		if count < 6 {
			result = result + string(r)
		} else {
			result = result + " "
			count = 0
		}
	}

	return result + "\n"
}
