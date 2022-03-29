package piscine

func IsPrintable(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] < 32 {
			return false
		}
	}
	return true
}
