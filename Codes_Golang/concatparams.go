package piscine

func ConcatParams(args []string) string {
	var str string
	i := 0
	for range args {
		i++
	}
	for ind, word := range args {
		str += word
		if ind != i-1 {
			str += "\n"
		}
	}
	return str
}
