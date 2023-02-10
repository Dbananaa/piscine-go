package piscine

func ReverseMenuIndex(menu []string) []string {
	reverseMenu := make([]string, len(menu))
	for i := len(menu) - 1; i >= 0; i-- {
		reverseMenu[len(menu)-i-1] = menu[i]
	}
	return reverseMenu
}
