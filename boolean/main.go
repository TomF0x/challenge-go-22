package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) int {
	if (nbr)%2 == 1 {
		return 1
	} else {
		return 2
	}
}

func main() {
	if isEven(len(os.Args[1:])) == 1 {
		printStr("I have an odd number of arguments")
	} else {
		printStr("I have an even number of arguments")
	}
}
