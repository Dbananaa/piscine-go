package piscine

func Rot14(str string) string {
	strNew := []rune(str)
	var result string

	for i := 0; i < len(strNew); i++ {
		if strNew[i] >= 'A' && strNew[i] <= 'L' {
			strNew[i] = strNew[i] + 'E'
		} else if strNew[i] >= 'M' && strNew[i] <= 'Z' {
			strNew[i] = strNew[i] - 'C'
		} else if strNew[i] >= 'a' && strNew[i] <= 'l' {
			strNew[i] = strNew[i] + 'E'
		} else if strNew[i] >= 'm' && strNew[i] <= 'z' {
			strNew[i] = strNew[i] - 'C'
		}
		result += string(strNew[i])
	}
	return result
}
