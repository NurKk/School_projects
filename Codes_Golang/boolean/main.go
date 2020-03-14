package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(str string) {
	arrayStr := []rune(str)
	length := 0

	for range arrayStr {
		length++
	}

	for i := 0; i < length; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	x := os.Args
	lengthOfArg := -1
	for range x {
		lengthOfArg++
	}
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if isEven(lengthOfArg) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
