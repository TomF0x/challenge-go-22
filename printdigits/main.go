package main

import "github.com/01-edu/z01"

func main() {
	digits := "0123456789\n"
	for _, letter := range digits {
		z01.PrintRune(letter)
	}
}
