package piscine

func MakeRange(min, max int) []int {
	var x []int
	y := 0
	if min < max {
		y = max - min
		x = make([]int, y)
		for i := 0; i < y; i++ {
			x[i] = min + i
		}
	} else {
		return nil
	}
	return x
}
