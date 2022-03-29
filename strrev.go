package piscine

func StrRev(s string) string {
	a := len(s)
	newword := ""
	for i := a - 1; i >= 0; i-- {
		newword += string(s[i])
	}
	return newword
}
