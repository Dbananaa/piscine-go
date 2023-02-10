package piscine

func ToUpper(s string) string {
	array := []rune(s)
	for j, i := range s {
		if i >= 'a' && i <= 'z' {
			array[j] -= 32
		}
	}
	return string(array)
}
