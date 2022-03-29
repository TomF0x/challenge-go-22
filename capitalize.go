package piscine

func Capitalize(s string) string {
	rep := ""
	tolower := ""
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] >= 65 && []rune(s)[k] <= 90 {
			tolower += string([]rune(s)[k] + 32)
		} else {
			tolower += string([]rune(s)[k])
		}
	}
	toup := true
	for k := 0; k < len([]rune(tolower)); k++ {
		if []rune(tolower)[k] >= 65 && []rune(tolower)[k] <= 90 || []rune(tolower)[k] >= '0' && []rune(tolower)[k] <= '9' {
			rep += string([]rune(tolower)[k])
			toup = false
		} else if []rune(tolower)[k] >= 97 && []rune(tolower)[k] <= 122 {
			if toup {
				rep += string([]rune(tolower)[k] - 32)
				toup = false
			} else {
				rep += string([]rune(tolower)[k])
				toup = false
			}
		} else {
			rep += string([]rune(tolower)[k])
			toup = true
		}
	}
	return rep
}
