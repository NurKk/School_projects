package piscine

func Rot14(str string) string {
	NewStr := ""
	for _, s := range str {
		if s >= 'A' && s <= 'Z' {
			if s+14 > 'Z' {
				y := ((s + 14) - 'Z') + 'A'
				NewStr += string(y - 1)
			} else {
				NewStr += string(s + 14)
			}
		} else if s >= 'a' && s <= 'z' {
			if s+14 > 'z' {
				x := ((s + 14) - 'z') + 'a'
				NewStr += string(x - 1)
			} else {
				NewStr += string(s + 14)
			}
		} else if s == ' ' {
			NewStr += string(' ')
		} else {
			NewStr += string(s)
		}
	}
	return NewStr
}
