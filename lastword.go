package piscine

// i = len - 1 -> 0

// s = "one two"
// i = 6 -> 0
// word = ""

// i = 6: word = "o" + "" = "o"
// i = 5: word = "w" + "o" = "wo"
// i = 4: word = "t" + "wo" = "two"
// i = 3: break
// result: word = "two"

// s = "one two   "
// word = "two"

func LastWord(s string) string {
	word := ""
	i := len(s) - 1

	// пройтись до первого не пробела
	for ; i >= 0 && s[i] == ' '; i-- {
	}

	// пройтись до следующего пробела
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			word = string(s[i]) + word
		} else {
			break
		}
	}
	return word + "\n"
}
