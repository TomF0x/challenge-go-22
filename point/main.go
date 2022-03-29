package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	str := [2]string{"x = ", ", y = "}
	for i := 0; i < 2; i++ {
		for _, letter := range str[i] {
			z01.PrintRune(letter)
		}
		if i == 0 {
			n := points.x
			finalnumber := 0
			lennumber := 0
			for nb := 1; nb < n; nb *= 10 {
				lennumber += 1
				finalnumber = nb
			}
			for nb := finalnumber; nb > 0; nb /= 10 {
				z01.PrintRune(rune((n / nb) + 48))
				n -= (n / nb) * nb
			}
		} else if i == 1 {
			n := points.y
			finalnumber := 0
			lennumber := 0
			for nb := 1; nb < n; nb *= 10 {
				lennumber += 1
				finalnumber = nb
			}
			for nb := finalnumber; nb > 0; nb /= 10 {
				z01.PrintRune(rune((n / nb) + 48))
				n -= (n / nb) * nb
			}
		}
	}
	z01.PrintRune('\n')
}
