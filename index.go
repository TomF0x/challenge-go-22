package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for count, letter := range s {
		if letter == rune(toFind[0]) {
			if s[count:count+len(toFind)] == toFind {
				return count
			}
		}
	}
	return -1
}
