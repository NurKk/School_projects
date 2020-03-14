package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	x := os.Args
	lenY := -1

	for range x {
		lenY++
	}
	for i := lenY; i > 0; i-- {
		for _, word := range x[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')

	}
}
