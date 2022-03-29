package piscine

func SortIntegerTable(table []int) {
	for k := 0; k < len(table)-1; k++ {
		i_min := k
		for i := k + 1; i < len(table); i++ {
			if table[i] < table[i_min] {
				i_min = i
			}
		}
		table[k], table[i_min] = table[i_min], table[k]
	}
}
