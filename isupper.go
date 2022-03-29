package piscine

func IsUpper(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] <= 'A' || []rune(s)[k] >= 'Z' {
			return false
		}
	}
	return true
}
