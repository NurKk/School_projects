package main

import (
	"github.com/01-edu/z01"
	// "fmt"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	y := 0
	length := 0
	strX := " ,"

	for i := 0; i < points.x; i = i {
		y = points.x % 10
		points.x /= 10
		strX += string(y + 48)
	}
	strX += " = x"

	for range strX {
		length++
	}
	for i := length - 1; i >= 0; i-- {

		z01.PrintRune(rune(strX[i]))
	}
	Last(points)
}

func Last(val *point) {
	strY := ""
	n := 0
	length := 0
	for i := 0; i < val.y; i = i {
		n = val.y % 10
		val.y /= 10
		strY += string(n + 48)
	}
	strY += " = y"
	for range strY {
		length++
	}
	for i := length - 1; i >= 0; i-- {
		z01.PrintRune(rune(strY[i]))
	}
	z01.PrintRune('\n')
}
