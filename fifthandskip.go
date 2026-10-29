package piscine

// "abcdefghijklmnopqrstuwxyz"

func FifthAndSkip(str string) string {
	runes := []rune(str) // конвертация

	if len(runes) == 0 {
		return "\n"
	}

	if len(runes) < 5 {
		return "Invalid Input\n"
	}

	count := 0   // счетчик символов
	result := "" // результат

	for _, r := range str {
		if r == ' ' { // пропускаю пробелы
			continue
		}

		count++ // считаю символы
		if count < 6 {
			result = result + string(r)
		} else {
			result = result + " " // замена каждого 6-го символа на пробел
			count = 0             // обнуление счетчика
		}
	}

	return result + "\n"
}
