package piscine

func CollatzCountdown(start int) int {
	result := 0
	if start <= 0 {
		return -1
	} else {
		for start > 1 {
			if start%2 != 0 {
				start = start*3 + 1
			} else {
				start = start / 2
			}
			result++
		}
	}
	return result
}
