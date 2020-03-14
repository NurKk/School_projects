package piscine

func IsLower(str string) bool {
	for _, word := range str {
		if word >= 'a' && word <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
