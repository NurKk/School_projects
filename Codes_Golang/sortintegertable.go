package piscine

func SortIntegerTable(table []int) {

	for i := 0; i < len(table)-1; i = i {
		if table[i] >= table[i+1] {
			table[i], table[i+1] = table[i+1], table[i]
			i = 0
		} else {
			i++
		}
	}
}
