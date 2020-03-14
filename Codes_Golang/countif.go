package piscine

func CountIf(f func(string) bool, tab []string) int {
	length := 0
	for _, word := range tab {
		if f(word) {
			length++
		}
	}
	if length > 0 {
		return length
	}
	return 0
}
