package piscine

import (
	"unicode"

	"github.com/01-edu/z01"
)

func printChars(bytes [10]byte) {
	for _, b := range bytes {
		r := rune(b)
		if unicode.IsGraphic(r) {
			z01.PrintRune(r)
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
