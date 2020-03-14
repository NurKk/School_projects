package piscine

func StrRev(s string) string {
	var y int
	var str string
	for y = range s {
		y++
	}
	for i := y - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}
