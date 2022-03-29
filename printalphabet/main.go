package main

import "github.com/01-edu/z01"

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz\n"
	for _, letter := range alpha {
		z01.PrintRune(letter)
	}
}
