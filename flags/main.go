package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0][:2] == "-h" || len(args[0]) > 5 && args[0][:6] == "--help" {
		stringhelp := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n"
		for _, letter := range stringhelp {
			z01.PrintRune(letter)
		}
	} else {
		toorder := false
		insert := ""
		numberofargs := 0
		for _, arg := range args {
			if len(arg) >= 2 {
				if "-o" == arg[:2] {
					numberofargs++
					toorder = true
				} else if len(arg) > 2 && "-i=" == arg[:3] {
					insert += arg[3:]
					numberofargs++
				} else if "--" == arg[:2] || "-" == arg[:1] {
					if arg[2:7] == "order" {
						numberofargs++
						toorder = true
					} else if len(arg) > 3 && arg[2:7] == "inser" {
						insert += arg[9:]
						numberofargs++
					}
				}
			}
		}
		args = args[numberofargs:]
		defaultstring := args[0] + insert
		if toorder {
			myletterlist := []rune{}
			for _, letter := range defaultstring {
				myletterlist = append(myletterlist, letter)
			}
			for k := 0; k < len(myletterlist)-1; k++ {
				i_min := k
				for i := k + 1; i < len(myletterlist); i++ {
					if myletterlist[i] < myletterlist[i_min] {
						i_min = i
					}
				}
				myletterlist[k], myletterlist[i_min] = myletterlist[i_min], myletterlist[k]
			}
			for _, letter := range myletterlist {
				z01.PrintRune(letter)
			}
		} else {
			for _, letter := range defaultstring {
				z01.PrintRune(letter)
			}
		}
		z01.PrintRune('\n')
	}
}
