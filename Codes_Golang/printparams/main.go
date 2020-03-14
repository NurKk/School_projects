package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	x := os.Args
	len := 0

	for i := range x {
		if i > 0 {
			y := []rune(os.Args[i])
			len = 0
			for range y {
				len++
			}
			for i := 0; i < len; i++ {
				z01.PrintRune(y[i])
			}
			z01.PrintRune('\n')

		}
	}
}
