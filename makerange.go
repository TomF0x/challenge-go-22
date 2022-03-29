package piscine

func MakeRange(min, max int) []int {
	if min < max {
		rep := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			rep[i] = i + min
		}
		return rep
	} else {
		return []int(nil)
	}
}
