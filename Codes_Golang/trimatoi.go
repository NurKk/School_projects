package piscine

func TrimAtoi(s string) int {
	var x int = 0
	var len int = 0
	var last int
	symbol := 1
	for _, word := range s {
		len = 0
		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {
				len++
				last++
			}
			x = x*10 + len
		}
		if word == '-' && last == 0 {
			symbol *= -1
		}
	}
	return x * symbol
}
