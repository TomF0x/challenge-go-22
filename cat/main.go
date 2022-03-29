package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				for _, letter := range "ERROR: open " {
					z01.PrintRune(letter)
				}
				for _, letter := range filename {
					z01.PrintRune(letter)
				}
				for _, letter := range ": no such file or directory" {
					z01.PrintRune(letter)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			arr := make([]byte, 100000)
			file.Read(arr)
			newbyte := []byte{}
			for _, bt := range arr {
				if bt != 0 {
					newbyte = append(newbyte, bt)
				}
			}
			os.Stdout.Write(newbyte)
		}
	} else {
		bytes, _ := os.Open(os.Stdin.Name())
		arr := make([]byte, 100000)
		bytes.Read(arr)
		newbyte := []byte{}
		for _, bt := range arr {
			if bt != 0 {
				newbyte = append(newbyte, bt)
			}
		}
		os.Stdout.Write(newbyte)
	}
}
