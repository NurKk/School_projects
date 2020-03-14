package piscine

func SplitWhiteSpaces(str string) []string {
	var array []string
	var lastString string
	lenInput := 0
	var NewLen int = 0
	num := 1
	for range str {
		lenInput++
	}
	for ind, word := range str {
		if word == ' ' || ind == lenInput-1 {
			// NewLen++
			num = 1
		} else if word != ' ' && num == 1 {
			NewLen++
			num = 0
		}
	}
	array = make([]string, NewLen)
	NewLen = 0
	bools := false
	for ind, word := range str {
		if word != 32 {
			lastString += string(word)
			bools = true
		}
		if word == ' ' && bools == true || ind == lenInput-1 && bools == true {
			array[NewLen] = lastString
			lastString = ""
			NewLen++
			bools = false
		}
	}
	return array
}
