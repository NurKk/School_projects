package piscine

func IsUpper(str string) bool {
	for _, word := range str {
		if word <= 'Z' && word >= 'A' {
			continue
		} else {
			return false
		}
	}
	return true
}
