package piscine

func Split(s, sep string) []string {
	rep := []string{}
	start := 0
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			rep = append(rep, s[start:i])
			start = i + len(sep)
		}
	}
	rep = append(rep, s[start:])
	return rep
}
