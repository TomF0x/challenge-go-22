package piscine

func ToLower(s string) string {
	stringtolower := ""
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] >= 65 && []rune(s)[k] <= 90 {
			stringtolower += string([]rune(s)[k] + 32)
		} else {
			stringtolower += string([]rune(s)[k])
		}
	}
	return stringtolower
}
