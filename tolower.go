package piscine

func ToLower(s string) string {
	array := []rune(s)
	for j, i := range s {
		if i >= 'A' && i <= 'Z' {
			array[j] += 32
		}
	}
	return string(array)
}
