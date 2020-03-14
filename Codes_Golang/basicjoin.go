package piscine

func BasicJoin(strs []string) string {
	x := []string(strs)
	// len := 0
	var str string
	for _, word := range x {
		str += word
	}
	return str
}
