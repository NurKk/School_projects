package piscine

func IsPrintable(str string) bool {
	x := []byte(str)
	// for _, word := range str {
	// 	if x >= 32  && x <= 127 {
	// 		continue
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true

	len := 0
	for range str {
		len++
	}

	for i := 0; i <= len-1; i++ {
		if x[i] >= 32 && x[i] <= 127 /*&& x[i] != 92 && x[i+1] != 110*/ {
			if x[i] == 92 && x[i+1] == 110 {
				return false
			} else {
				continue
			}
		} else {
			return false
		}

	}
	return true
}
