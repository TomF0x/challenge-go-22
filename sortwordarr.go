package piscine

func SortWordArr(a []string) {
	for k := 0; k < len(a)-1; k++ {
		i_min := k
		for i := k + 1; i < len(a); i++ {
			if a[i] < a[i_min] {
				i_min = i
			}
		}
		a[k], a[i_min] = a[i_min], a[k]
	}
}
