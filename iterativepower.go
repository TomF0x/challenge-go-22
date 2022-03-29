package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	}
	rep := nb
	for i := 0; i < power-1; i++ {
		rep *= nb
	}
	return rep
}
