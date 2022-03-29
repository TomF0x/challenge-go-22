package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for k := nb; k < nb*2; k++ {
		tempbool := true
		if nb > 100000 {
			for i := 2; i < nb/10; i++ {
				if k%i == 0 {
					tempbool = false
				}
			}
		} else {
			for i := 2; i < nb; i++ {
				if k%i == 0 {
					tempbool = false
				}
			}
		}
		if tempbool {
			return k
		}
	}
	return 0
}
