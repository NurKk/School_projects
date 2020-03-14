package piscine

func Capitalize(s string) string {
	x := []rune(s)
	len := 0
	for range s {
		len++
	}
	if len > 1 {
		for i := 1; i < len; i++ {
			if x[0] >= 97 && x[0] <= 122 {
				x[0] -= 32
			}
			if x[i-1] >= 32 && x[i-1] <= 47 || x[i-1] >= 58 && x[i-1] <= 64 || x[i-1] >= 91 && x[i-1] <= 96 || x[i-1] >= 123 && x[i-1] <= 127 {
				if x[i] >= 97 && x[i] <= 122 {
					x[i] -= 32
				}
			} else {
				if x[i] >= 65 && x[i] <= 90 {
					x[i] += 32
				}
			}
		}
	} else {
		if x[0] >= 97 && x[0] <= 122 {
			x[0] -= 32
		}
	}
	str := string(x)
	return str
}
