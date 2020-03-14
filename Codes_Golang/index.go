package piscine

func Index(s string, toFind string) int {
	x, y := 0, 0
	for range s {
		x++
	}
	for range toFind {
		y++
	}
	for i := 0; i <= x; i++ {
		if i+y > x {
			break
		} else if toFind == s[i:i+y] {
			return i
		}
	}
	return -1
}
