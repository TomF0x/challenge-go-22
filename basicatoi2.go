package piscine

func BasicAtoi2(s string) int {
	number := 0
	lenbumber := len(s)
	grandeur := 1
	for i := lenbumber - 1; i >= 0; i-- {
		if []byte(s)[i] > 57 || []byte(s)[i] < 48 {
			return 0
		} else {
			number += int([]byte(s)[i]-48) * grandeur
			grandeur *= 10
		}
	}
	return number
}
