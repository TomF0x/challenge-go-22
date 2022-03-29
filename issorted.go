package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	croissant, decroissant := true, true
	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) < 0 {
			decroissant = false
		}
	}

	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			croissant = false
		}
	}
	return croissant || decroissant
}
