package piscine

func Atoi(s string) int {
	number := 0
	isneg := false
	lenbumber := len(s)
	if lenbumber == 0 {
		return 0
	}
	grandeur := 1
	if []byte(s)[0] == 45 {
		isneg = true
		tempstirng := []byte(s)
		tempstirng[0] = '0'
		s = string(tempstirng)
	} else if []byte(s)[0] == 43 {
		tempstirng := []byte(s)
		tempstirng[0] = '0'
		s = string(tempstirng)
	}
	for i := lenbumber - 1; i >= 0; i-- {
		if []byte(s)[i] > 57 || []byte(s)[i] < 48 {
			return 0
		} else {
			number += int([]byte(s)[i]-48) * grandeur
			grandeur *= 10
		}
	}
	if isneg {
		number *= -1
	}
	return number
}
