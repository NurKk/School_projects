package piscine

func AlphaCount(str string) int {
	count := 0
	for _, Rune := range str {
		if Rune >= 'a' && Rune <= 'z' || Rune >= 'A' && Rune <= 'Z' {
			count++
		}
	}
	return count
}
