package piscine

func IsLower(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] < 'a' || []rune(s)[k] > 'z' {
			return false
		}
	}
	return true
}
