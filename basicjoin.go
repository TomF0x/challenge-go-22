package piscine

func BasicJoin(elems []string) string {
	join := ""
	for _, word := range elems {
		join += word
	}
	return join
}
