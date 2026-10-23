package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	for i := 0; i < 4; i++ {
		convert16(arr[i])
		if i < 3 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
	}
	for i := 4; i < 8; i++ {
		convert16(arr[i])
		if i < 7 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
	}
	for i := 8; i < 10; i++ {
		convert16(arr[i])
		if i < 9 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
	}
	for _, b := range arr {
		if b >= 33 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')

}

// d1 = 0 -> d1 + '0' = 0 + 48 = 48 -> '0'
// d1 = 1 -> d1 + '0' = 1 + 48 = 49 -> '1'
// d1 = 3 -> d1 + '0' = 3 + 48 = 51 -> '3'

// d1 = 10 -> d1 - 10 + 'a' = 10 - 10 + 97 = 97 -> 'a'
// d1 = 11 -> d1 - 10 + 'a' = 11 - 10 + 97 = 98 -> 'b'

// 21 = 2 * 10 + 1 * 1
// 756 = 7 * 100 + 5 * 10 + 6 * 1

// 51[16] = 5 * 16 + 1 * 1 = 81[10]
// 81 / 16 = 5; 81 % 16 = 1

func convert16(b byte) {
	d1 := b / 16
	d2 := b % 16

	if d1 < 10 {
		z01.PrintRune(rune(d1) + '0')
	} else {
		z01.PrintRune(rune(d1) - 10 + 'a')
	}

	if d2 < 10 {
		z01.PrintRune(rune(d2) + '0')
	} else {
		z01.PrintRune(rune(d2) - 10 + 'a')
	}
}
