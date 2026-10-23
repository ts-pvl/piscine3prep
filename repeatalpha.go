package piscine

// s = "abc"
// result = ""

// r = 'a' = 97
// i = 97 -> 97 result = "a"

// r = 'b' = 98
// i = 97 -> 98 (1 time) result = "abb"

// r = 'c' = 99
// i = 97 -> 99 (2 time) result = "abbccc"

func RepeatAlpha(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			for i := 'a'; i < r; i++ {
				result = result + string(r)
			}
		}
		if r >= 'A' && r <= 'Z' {
			for i := 'A'; i < r; i++ {
				result = result + string(r)
			}
		}
		result = result + string(r)
	}
	return result
}
