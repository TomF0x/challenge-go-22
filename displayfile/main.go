package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else {
		file, _ := os.Open(os.Args[1])
		arr := make([]byte, 14)
		file.Read(arr)
		fmt.Println(string(arr))
	}
}
