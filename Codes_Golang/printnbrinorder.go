package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	x := []rune{}
	y := 0
	len := 0
	if n == 0 {
		x = append(x, rune(n))
	}
	for n != 0 {
		y = n % 10
		x = append(x, rune(y))
		n /= 10
	}
	for range x {
		len++
	}
	for i := 0; i < len-1; i++ {
		for j := i; j < len; j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	for m := 0; m < len; m++ {
		z01.PrintRune(rune(x[m]) + 48)
	}
}
