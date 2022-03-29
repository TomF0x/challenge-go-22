package piscine

func AppendRange(min, max int) []int {
	rep := []int{}
	if min == 0 && max == 0 || max < min {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		rep = append(rep, i)
	}
	return rep
}
