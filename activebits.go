package piscine

func ActiveBits(n int) int {
	count := 0
	binary := DecToBin(n)
	for _, v := range binary {
		if v == '1' {
			count++
		}
	}
	return count
}

func DecToBin(n int) string {
	result := ""

	for n > 0 {
		temp := n % 2
		result += string(rune(temp + '0'))
		n /= 2
	}
	return result
}
