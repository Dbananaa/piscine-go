package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	array := make([]int, max-min)

	for i := 0; i < max-min; i++ {
		array[i] = min + i
	}
	return array
}
