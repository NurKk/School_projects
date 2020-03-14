package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	len := 0
	x := os.Args
	for range x {
		len++
	}
	for i := 0; i < len-1; i = i {
		if x[i] > x[i+1] {
			x[i], x[i+1] = x[i+1], x[i]
			i = 0
		} else {
			i++
		}
	}
	for i := 0; i < len-1; i++ {
		for _, word := range x[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
