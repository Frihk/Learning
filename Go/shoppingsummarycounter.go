package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	current := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			current += string(str[i])
		} else {
			result[current]++
			current = ""
		}
	}
	result[current]++
	return result
}
