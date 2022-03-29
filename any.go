package piscine

func Any(f func(string) bool, a []string) bool {
	for _, word := range a {
		if f(word) == true {
			return true
		}
	}
	return false
}
