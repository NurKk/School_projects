package piscine

func RecursivePower(nb int, power int) int {

	if power >= 0 {

		if power == 0 {
			return 1
		} else {

			return nb * RecursivePower(nb, power-1)

		}
	} else {
		return 0
	}

	return nb

}
