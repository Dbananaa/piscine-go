package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	c := len(s)
	if n <= c && n > 0 {
		return a[n-1]
	} else {
		return 0
	}
}
