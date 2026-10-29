package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	runes1 := []rune(s1)
	s2 := os.Args[2]

	if len(s1) == 0 {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	i1 := 0
	for _, r := range s2 {
		if i1 == len(runes1) {
			z01.PrintRune('1')
			z01.PrintRune('\n')
			return
		}
		if r == runes1[i1] {
			i1++
		}
	}
	if i1 == len(runes1) {
		z01.PrintRune('1')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}
