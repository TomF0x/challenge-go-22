package piscine

func ForEach(f func(int), a []int) {
	for _, b := range a {
		f(b)
	}
}
