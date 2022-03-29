package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for k := 0; k < len(args)-1; k++ {
		i_min := k
		for i := k + 1; i < len(args); i++ {
			if args[i] < args[i_min] {
				i_min = i
			}
		}
		args[k], args[i_min] = args[i_min], args[k]
	}
	for _, arg := range args {
		for _, letter := range arg {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
