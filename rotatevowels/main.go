package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
	} else {
		newLine := ""
		voy := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
		voyels := []rune{}
		for _, k := range args {
			newLine += k + " "
		}
		for _, i := range newLine {
			for _, x := range voy {
				if i == x {
					voyels = append(voyels, i)
				}
			}
		}
		for i := 0; i < (len(voyels) / 2); i++ {
			voyels[i], voyels[len(voyels)-1-i] = voyels[len(voyels)-1-i], voyels[i]
		}
		voyelsdone := 0
		newlinebyte := []byte(newLine)
		for number, i := range newlinebyte {
			for _, x := range voy {
				if rune(i) == x {
					newlinebyte[number] = byte(voyels[voyelsdone])
					voyelsdone++
				}
			}
		}
		for _, letter := range newlinebyte {
			z01.PrintRune(rune(letter))
		}
		z01.PrintRune('\n')
	}
}
