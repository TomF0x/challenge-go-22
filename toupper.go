package piscine

func ToUpper(s string) string {
	stringerupper := ""
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] >= 97 && []rune(s)[k] <= 122 {
			stringerupper += string([]rune(s)[k] - 32)
		} else {
			stringerupper += string([]rune(s)[k])
		}
	}
	return stringerupper
}
