package piscine

func ToLower(s string) string {
	x := []byte(s)
	var str string

	for i := 0; i < len(s); i++ {
		if x[i] >= 65 && x[i] <= 90 {
			str += string(x[i] + 32)
		} else if x[i] >= 32 && x[i] <= 127 {
			str += string(x[i])
		}
	}
	return str
}
