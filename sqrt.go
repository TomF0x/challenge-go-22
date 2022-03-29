package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for k := 0; k < (nb/2)+1; k++ {
		rep := k
		rep *= rep
		if rep == nb {
			return k
		}
	}
	return 0
}
