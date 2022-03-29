package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	value := []int{0, 0, 0}
	for i := 0; i < 999; i++ {
		if value[2] > value[1] && value[1] > value[0] {
			z01.PrintRune(rune(value[0] + 48))
			z01.PrintRune(rune(value[1] + 48))
			z01.PrintRune(rune(value[2] + 48))
			if value[0]+value[1]+value[2] != 24 {
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
		value[2]++
		if value[2] == 10 {
			value[2] = 0
			value[1]++
		}
		if value[1] == 10 {
			value[1] = 0
			value[0]++
		}
	}
	z01.PrintRune('\n')
}
