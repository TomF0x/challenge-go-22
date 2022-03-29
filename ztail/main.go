package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	chr := 0
	haveerror := false
	if len(args) >= 1 {
		if args[0] == "-c" {
			s := args[1]
			number := 0
			lenbumber := len(s)
			grandeur := 1
			for i := lenbumber - 1; i >= 0; i-- {
				number += int([]byte(s)[i]-48) * grandeur
				grandeur *= 10
			}
			chr = number
			args = args[2:]
		}
	}
	if len(args) == 1 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Printf("open %v: no such file or directory\n", args[0])
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
		if len(newbyte)+1-chr < 0 {
			os.Stdout.Write(newbyte)
		} else {
			os.Stdout.Write(newbyte[len(newbyte)+1-chr:])
		}
	} else if len(args) > 1 {
		for number, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Printf("open %v: no such file or directory\n", filename)
				haveerror = true
			} else {
				if number != 0 {
					fmt.Println("")
				}
				fmt.Printf("==> %v <==\n", filename)
				arr := make([]byte, 100000)
				file.Read(arr)
				newbyte := []byte{}
				for _, bt := range arr {
					if bt != 0 {
						newbyte = append(newbyte, bt)
					}
				}
				if len(newbyte)+1-chr < 0 {
					os.Stdout.Write(newbyte)
				} else {
					os.Stdout.Write(newbyte[len(newbyte)-chr:])
				}
			}
		}
	}
	if haveerror {
		os.Exit(1)
	}
}
