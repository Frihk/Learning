package piscine

func ReverseMenuIndex(menu []string) []string {
	nmenu := make([]string, len(menu))
	for i := len(menu) - 1; i >= 0; i-- {
		nmenu[len(menu)-1-i] = menu[i]
	}
	return nmenu
}
