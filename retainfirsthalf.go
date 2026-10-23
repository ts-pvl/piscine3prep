package piscine

// str: 1234
// polova: 12

// i = 0 -> 2 (4/2)

// -----------------

// str: abcdef
// i = 0 -> 3

// i = 0 -> polovina: a
// i = 1 -> polovina: ab
// i = 2 -> polovina: abc
// i = 3 -> stop!

// -------------------

// 5 / 2 = 2 ~ 4 / 2 = 2

func RetainFirstHalf(str string) string {
	if str == "" {
		return ""
	}

	if len(str) == 1 {
		return str
	}

	polovina := ""
	for i := 0; i < len(str)/2; i++ {
		polovina = polovina + string(str[i])
	}
	return polovina
}
