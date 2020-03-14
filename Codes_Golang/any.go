package piscine

func Any(n func(string) bool, arr []string) bool {
	length := 0
	for _, word := range arr {
		if n(word) == true {
			length++
		}
	}
	if length > 0 {
		return true
	}
	return false
}
