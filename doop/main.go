package main

import (
	"os"
)

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

func IsNumeric(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] < 48 || []rune(s)[k] > 57 {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	for k := 0; k < len(args); k++ {
		if !(IsNumeric(args[k]) || args[k][0] == '-' && IsNumeric(args[k][1:]) || args[k] == "+" || args[k] == "*" || args[k] == "/" || args[k] == "-" || args[k] == "%") {
			os.Exit(0)
		}
	}
	number := Atoi(args[0])
	if number >= 9223372036854775807 || number <= -9223372036854775807 {
		os.Exit(0)
	}
	for k := 1; k < len(args)-1; k++ {
		if args[k] == "+" {
			number += Atoi(args[k+1])
		} else if args[k] == "*" {
			number *= Atoi(args[k+1])
		} else if args[k] == "-" {
			number -= Atoi(args[k+1])
		} else if args[k] == "/" && Atoi(args[k+1]) != 0 {
			number /= Atoi(args[k+1])
		} else if args[k] == "%" && Atoi(args[k+1]) != 0 {
			number %= Atoi(args[k+1])
		} else {
			if args[k] == "/" {
				os.Stdout.WriteString("No division by 0\n")
			} else if args[k] == "%" {
				os.Stdout.WriteString("No modulo by 0\n")
			}
			os.Exit(0)
		}
		if number < 0 {
			os.Stdout.WriteString("-")
			number = number * -1
		}
		if number == 1 {
			os.Stdout.WriteString("1")
		}
		if number == 0 {
			os.Stdout.WriteString("0")
		}
		n := number
		finalnumber := 0
		lennumber := 0
		for nb := 1; nb < n; nb *= 10 {
			lennumber += 1
			finalnumber = nb
		}
		test := []byte{}
		for nb := finalnumber; nb > 0; nb /= 10 {
			test = append(test, byte(rune((n/nb)+48)))
			n -= (n / nb) * nb
		}
		os.Stdout.Write(test)
		os.Stdout.WriteString("\n")
	}
}
