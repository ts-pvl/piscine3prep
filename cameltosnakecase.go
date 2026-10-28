package piscine

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	if isInvalidCamelCase(s) {
		return s
	}

	// convert
	result := ""
	for i, r := range s {
		if i != 0 && isUpper(r) {
			result = result + "_"
		}
		result = result + string(r)
	}
	return result
}

// true - если не валидный camel case
func isInvalidCamelCase(s string) bool {
	runes := []rune(s)

	// проверка 1: больщая буква в конце
	lastRune := runes[len(runes)-1]
	if isUpper(lastRune) {
		return true
	}

	// проверка 2: 2 большие буквы рядом
	for i := 0; i < len(runes)-1; i++ {
		if isUpper(runes[i]) && isUpper(runes[i+1]) {
			return true
		}
	}

	for _, r := range runes {
		if r >= '0' && r <= '9' {
			return true
		}
		if r == '.' || r == ',' || r == '_' {
			return true
		}
	}

	return false
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
