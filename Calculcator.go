package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 3 {

		var result int
		oper := args[1]

		n1, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("0")
			return
		}

		n2, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("0"
			return
		}

		if oper == "+" {
			if n1 >= 0 && n2 >= 0 && n1+n2 < 0 {
				fmt.Println("0")
				return
			}
			result = n1 + n2
		} else if oper == "-" {
			if n1 < 0 && n2 < 0 && n1+n2 > 0 {
				fmt.Println("0")
				return
			}
			result = n1 - n2
		} else if oper == "*" {
			if n1 > 0 && n2 > 0 && n1+n2 < 0 {
				fmt.Println("0")
				return
			}
			result = n1 * n2
		} else if oper == "/" {
			if n2 == 0 {
				fmt.Println("no division by 0")
				return
			}
			result = n1 / n2
		} else if oper == "%" {
			if n2 == 0 {
				fmt.Println("no modulo by 0")
				return
			}
			result = n1 % n2
		}
		fmt.Print(result)

	}
	fmt.Print("\n")
}
