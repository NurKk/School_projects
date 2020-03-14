package piscine

func IsAlpha(str string) bool {
	for _, word := range str {
		if word >= 'A' && word <= 'Z' || word >= 'a' && word <= 'z' || word >= '0' && word <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
