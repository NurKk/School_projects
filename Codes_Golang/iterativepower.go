package piscine

func IterativePower(nb int, power int) int {
	x := 1
	if power >= 0 {
		for i := 1; i <= power; i++ {

			x *= nb

		}

	} else {
		return 0
	}
	return x
}
