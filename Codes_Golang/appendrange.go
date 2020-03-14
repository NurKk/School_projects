package piscine

func AppendRange(min, max int) []int {
	x := []int{}
	if min < max {
		for i := min; i < max; i++ {
			x = append(x, i)
		}
	} else {
		return nil
	}
	return x
}
