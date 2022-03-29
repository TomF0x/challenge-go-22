package piscine

func IsNumeric(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] < 48 || []rune(s)[k] > 57 {
			return false
		}
	}
	return true
}
