package piscine

func AlphaCount(s string) int {
	array := []rune(s)
	count := 0
	for i := 0; i < len(array); i++ {
		if (array[i] >= 'a' && array[i] <= 'z') || (array[i] >= 65 && array[i] <= 90) {
			count++
		}
	}
	return count
}
