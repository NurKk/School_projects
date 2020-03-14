package piscine

func Join(strs []string, sep string) string {
	var str string
	len := 0
	for range strs {
		len++
	}
	for ind, word := range strs {
		str += word
		if len-1 != ind {
			str += sep
		}
	}
	return str
}
