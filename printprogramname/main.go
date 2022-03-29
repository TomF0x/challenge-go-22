package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	index := 0
	for number, letter := range name {
		if string(letter) == "./" {
			index = number
		}
	}
	name = name[index+2:]
	for _, k := range name {
		z01.PrintRune(k)
	}
	z01.PrintRune('\n')
}
