package piscine

func CountIf(f func(string) bool, tab []string) int {
	rep := 0
	for _, word := range tab {
		if f(word) == true {
			rep += 1
		}
	}
	return rep
}
