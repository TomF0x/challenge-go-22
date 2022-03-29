package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 1000 {
		return 0
	}
	rep := 1
	for i := 1; i < nb+1; i++ {
		rep *= i
	}
	return rep
}
