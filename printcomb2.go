package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	value := []int{0, 0, 0, 0}
	for i := 0; i < 9999; i++ {
		if ((value[0] * 10) + value[1]) < ((value[2] * 10) + value[3]) {
			z01.PrintRune(rune(value[0] + 48))
			z01.PrintRune(rune(value[1] + 48))
			z01.PrintRune(rune(32))
			z01.PrintRune(rune(value[2] + 48))
			z01.PrintRune(rune(value[3] + 48))
			if (value[0]*1000)+(value[1]*100)+((value[2]*10)+value[3]) != 9899 {
				z01.PrintRune(rune(44))
				z01.PrintRune(rune(32))
			}
		}
		value[3] += 1
		if value[3] == 10 {
			value[3] = 0
			value[2] += 1
		}
		if value[2] == 10 {
			value[2] = 0
			value[1] += 1
		}
		if value[1] == 10 {
			value[1] = 0
			value[0] += 1
		}
	}
	z01.PrintRune('\n')
}
