package piscine

func ForEach(f func(int), arr []int) {
	for _, word := range arr {
		f(word)
	}
}
