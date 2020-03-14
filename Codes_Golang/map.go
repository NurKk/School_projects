package piscine

func Map(f func(int) bool, arr []int) []bool {
	length := 0
	for range arr {
		length++
	}
	var x []bool = make([]bool, length)
	for i, word := range arr {
		x[i] = f(word)
	}
	return x
}
