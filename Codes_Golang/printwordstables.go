package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {

	y := 0

	for range table {
		y++
	}
	for i := 0; i < y; i++ {
		for _, word := range table[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
