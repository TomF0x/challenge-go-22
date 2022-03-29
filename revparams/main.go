package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		for _, letter := range args[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
