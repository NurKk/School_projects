package piscine

func Abort(a, b, c, d, e int) int {
	x := []int{a, b, c, d, e}
	length := 0
	for range x {
		length++
	}
	for i := 0; i < length-1; i = i {
		if x[i] > x[i+1] {
			x[i], x[i+1] = x[i+1], x[i]
			i = 0
		} else {
			i++
		}
	}
	return x[2]
}
