package piscine

func IsNumeric(str string) bool {
	for _, word := range str {
		if word >= '0' && word <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
