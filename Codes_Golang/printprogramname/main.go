package main

import (
	"github.com/01-edu/z01"
	"os"
)

 
func main() {

	x := []rune(os.Args[0])

	for i := range x {
		z01.PrintRune(x[i])
	}
	z01.PrintRune('\n')

}

