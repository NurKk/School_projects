package main

import (
	"github.com/01-edu/z01"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	x := os.Args
	length := 0
	for range x {
		length++
	}
	if length != 1 {
		file, err := ioutil.ReadFile(x[1])
		if err != nil {
			for _, word := range err.Error() {
				z01.PrintRune(word)
			}
			z01.PrintRune('\n')
		}
		if length == 2 {
			str := string(file)
			for _, word := range str {
				z01.PrintRune(word)
			}
		} else if length == 3 {
			fileTwo, _ := ioutil.ReadFile(os.Args[2])
			str := string(file)
			strTwo := string(fileTwo)
			for _, word := range str {
				z01.PrintRune(word)
			}
			for _, word := range strTwo {
				z01.PrintRune(word)
			}
		}
	} else {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {

		}
	}
}
