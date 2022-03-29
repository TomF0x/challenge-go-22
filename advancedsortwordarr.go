package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for k := 0; k < len(a)-1; k++ {
		i_min := k
		for i := k + 1; i < len(a); i++ {
			if f(a[i], a[i_min]) == -1 {
				i_min = i
			}
		}
		a[k], a[i_min] = a[i_min], a[k]
	}
}
