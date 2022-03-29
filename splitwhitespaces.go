package piscine

func SplitWhiteSpaces(s string) []string {
	start := 0
	rep := []string{}
	for loop, myrune := range s {
		if myrune == ' ' {
			rep = append(rep, s[start:loop])
			start = loop + 1
		}
		if loop == len(s)-1 {
			rep = append(rep, s[start:loop+1])
			start = loop + 1
		}
	}
	newrep := []string{}
	for _, word := range rep {
		if word != "" {
			newrep = append(newrep, word)
		}
	}
	return newrep
}
