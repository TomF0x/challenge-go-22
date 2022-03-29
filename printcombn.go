package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		a := 0
		for k := 0; k <= 9; k++ {
			a = k
			z01.PrintRune(rune(a) + 48)
			if a == 9 {
			} else {
				z01.PrintRune(rune(44))
				z01.PrintRune(rune(32))
			}
		}
		z01.PrintRune(rune('\n'))
	} else {
		allvar := append([]int{})
		for i := 0; i < n; i++ {
			allvar = append(allvar, 0)
		}
		grandeur := 1
		for i := 0; i < n; i++ {
			grandeur *= 10
		}
		firstprint := true
		if n == 9 {
			grandeur = 123456790
		}
		for i := 0; i < grandeur; i++ {
			listbool := append([]bool{})
			ntem := n - 1
			for i := 0; 0 < ntem; i++ {
				listbool = append(listbool, allvar[ntem-1] < allvar[ntem])
				ntem -= 1
			}
			passed := false
			for _, booltocheck := range listbool {
				if booltocheck {
					passed = true
				} else {
					passed = false
					break
				}
			}
			if passed {
				if firstprint {
					firstprint = false
				} else {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
				thirdntem := 0
				for i := 0; n > thirdntem; i++ {
					z01.PrintRune(rune(allvar[thirdntem] + 48))
					thirdntem += 1
				}
			}
			secondntem := n - 1
			for i := 0; 0 < secondntem; i++ {
				if allvar[secondntem] == 9 {
					allvar[secondntem] = 0
					allvar[secondntem-1] += 1
				}
				secondntem -= 1
			}
			allvar[n-1] += 1
		}
		z01.PrintRune(rune('\n'))
	}
}
