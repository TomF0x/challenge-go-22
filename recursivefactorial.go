package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb == 1 {
		return 1
	} else if nb > 1 {
		return nb * IterativeFactorial(nb-1)
	}
	return 0
}
