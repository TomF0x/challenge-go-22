package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	bignumber := false
	if n == 1 {
		z01.PrintRune('1')
	}
	if n == 0 {
		z01.PrintRune(48)
	}
	if n < 0 {
		z01.PrintRune(45)
		if n == -9223372036854775808 {
			bignumber = true
			n = 9223372036854775807
		} else {
			n = n * -1
		}
	}
	lennumber := 0
	finalnumber := 0
	if n > 1000000000000000000 {
		lennumber = 19
		finalnumber = 1000000000000000000
	} else {
		for nb := 1; nb < n; nb *= 10 {
			lennumber += 1
			finalnumber = nb
		}
	}
	for nb := finalnumber; nb > 0; nb /= 10 {
		if bignumber && n < 9 {
			n += 1
			bignumber = false
		}
		z01.PrintRune(rune((n / nb) + 48))
		n -= (n / nb) * nb
	}
}
