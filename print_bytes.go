package piscine

import "github.com/01-edu/z01"

func printBytes(bytes [10]byte) {
	for i := 0; i < 4; i++ {
		convert16(bytes[i])
		if i < 3 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
	}
	for i := 4; i < 8; i++ {
		convert16(bytes[i])
		if i < 7 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
	}
	convert16(bytes[8])
	z01.PrintRune(' ')
	convert16(bytes[9])
	z01.PrintRune('\n')
}
