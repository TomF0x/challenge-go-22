package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(48)
	} else {
		intarr := []int{}
		bignumber := false
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
			intarr = append(intarr, n/nb)
			n -= (n / nb) * nb
		}
		for k := 0; k < len(intarr)-1; k++ {
			i_min := k
			for i := k + 1; i < len(intarr); i++ {
				if intarr[i] < intarr[i_min] {
					i_min = i
				}
			}
			intarr[k], intarr[i_min] = intarr[i_min], intarr[k]
		}
		for _, nb := range intarr {
			z01.PrintRune(rune(nb + 48))
		}
	}
}
