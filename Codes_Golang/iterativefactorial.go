package piscine

func IterativeFactorial(nb int) int {
	x := 0
	if nb >= 0 && nb <= 20 {
		x = 1
		for i := 1; i < nb; i++ {
			x = x * (i + 1)
		}
	} else {
		return 0
	}
	return x
}
