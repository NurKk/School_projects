package piscine

func NRune(s string, n int) rune {
	var num rune
	var ind int
	var count int
	var Rune rune

	for ind, Rune = range s {
		count++
	}
	if count >= n && n > 0 {
		for ind, Rune = range s {
			num = Rune
			if ind+1 == n {
				return num
			}
		}
	}
	return num
}
