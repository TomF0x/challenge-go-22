package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	var number []int
	var temp int
	rep := ""
	for _, nb := range nbr {
		for i, p := range baseFrom {
			if p == nb {
				number = append(number, i)
			}
		}
	}
	temp = number[0]
	for x := 1; x < len(number); x++ {
		temp *= len(baseFrom)
		temp += number[x]
	}
	var s []int
	for i := 0; temp > 0; i++ {
		s = append(s, temp%len(baseTo))
		temp /= len(baseTo)
	}

	for x := len(s) - 1; x >= 0; x-- {
		rep += string(baseTo[s[x]])
	}
	return rep
}
