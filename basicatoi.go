package piscine

func BasicAtoi(s string) int {
	number := 0
	lenbumber := len(s)
	grandeur := 1
	for i := lenbumber - 1; i >= 0; i-- {
		number += int([]byte(s)[i]-48) * grandeur
		grandeur *= 10
	}
	return number
}
