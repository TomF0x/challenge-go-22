package piscine

func IsAlpha(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] >= 32 && []rune(s)[k] <= 47 || []rune(s)[k] >= 58 && []rune(s)[k] <= 64 {
			return false
		}
	}
	return true
}
