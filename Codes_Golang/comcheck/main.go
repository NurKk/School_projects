package main

import (
	"fmt"
	"os"
)

func main() {
	x := os.Args[1:]
	for _, word := range x {
		if word == "01" || word == "galaxy" || word == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
