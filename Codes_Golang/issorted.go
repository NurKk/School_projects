package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	tablen := -1
	lenMinus := 0
	lengthO := 0
	for range tab {
		tablen++
	}
	for i := 0; i < tablen; i++ {
		if f(tab[i], tab[i+1]) < 0 {
			lenMinus++
		} else if f(tab[i], tab[i+1]) > 0 {
			length++
		} else if f(tab[i], tab[i+1]) == 0 {
			lengthO++
		}
	}
	if length == tablen || lenMinus == tablen {
		return true
	} else if length != tablen && lengthO != tablen || lenMinus != tablen && lengthO != tablen {
		return false
	}
	return true
}
