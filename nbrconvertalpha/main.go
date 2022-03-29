package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		toup := false
		if args[0] == "--upper" {
			args = args[1:]
			toup = true
		}
		arryint := []int{}
		for _, mynumber := range args {
			number := 0
			lenbumber := len(mynumber)
			grandeur := 1
			for i := lenbumber - 1; i >= 0; i-- {
				number += int([]byte(mynumber)[i]-48) * grandeur
				grandeur *= 10
			}
			arryint = append(arryint, number)
		}
		if toup {
			for _, nb := range arryint {
				z01.PrintRune(rune(nb + 'A' - 1))
			}
		} else {
			for _, nb := range arryint {
				if nb < 26 {
					z01.PrintRune(rune(nb + 'a' - 1))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
