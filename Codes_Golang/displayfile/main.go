package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	x := os.Args
	length := 0
	for range x {
		length++
	}
	if length > 2 {
		fmt.Println("Too many arguments")

	} else if length == 1 {
		fmt.Println("File name missing")
	} else {
		file, _ := ioutil.ReadFile(os.Args[1])

		fmt.Println(string(file))
	}
}
