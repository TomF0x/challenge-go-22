package piscine

func Join(strs []string, sep string) string {
	rep := ""
	for _, word := range strs {
		rep += word
		rep += sep
	}
	rep = rep[:len(rep)-len(sep)]
	return rep
}
