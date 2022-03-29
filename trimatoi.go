package piscine

func TrimAtoi(s string) int {
	neword := ""
	for _, letter := range s {
		if letter == '-' && neword == "" || letter == '+' && neword == "" {
			neword += string(letter)
		}
		if letter >= '0' && letter <= '9' {
			neword += string(letter)
		}
	}
	number := 0
	isneg := false
	lenbumber := len(neword)
	if lenbumber == 0 {
		return 0
	}
	grandeur := 1
	if []byte(neword)[0] == 45 {
		isneg = true
		tempstirng := []byte(neword)
		tempstirng[0] = '0'
		neword = string(tempstirng)
	} else if []byte(neword)[0] == 43 {
		tempstirng := []byte(neword)
		tempstirng[0] = '0'
		neword = string(tempstirng)
	}
	for i := lenbumber - 1; i >= 0; i-- {
		if []byte(neword)[i] > 57 || []byte(neword)[i] < 48 {
			return 0
		} else {
			number += int([]byte(neword)[i]-48) * grandeur
			grandeur *= 10
		}
	}
	if isneg {
		number *= -1
	}
	return number
}
