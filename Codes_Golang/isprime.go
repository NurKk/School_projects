package piscine

func IsPrime(nb int) bool {
	x := 0

	if nb > 3 {
		for i := 1; i <= nb; i++ {
			if nb%i == 0 {
				x++
			}
			if x > 2 {
				return false
			}

		}
		return true

	} else if nb <= 1 {
		return false
	}

	return true
}
