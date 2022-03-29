package piscine

func Map(f func(int) bool, a []int) []bool {
	rep := []bool{}
	for _, nbr := range a {
		rep = append(rep, f(nbr))
	}
	return rep
}
