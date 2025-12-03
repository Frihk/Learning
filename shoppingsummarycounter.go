package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	current := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' && str[i] != '\t' && str[i] != '\n' {
			current += string(str[i])
		} else {
			if current != "" {
				result[current]++
				current = ""
			}
		}
	}
	if current != "" {
		result[current]++
	}
	return result
}
