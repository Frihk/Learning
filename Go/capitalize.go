package piscine

func Capitalize(s string) string {
	start := true
	result := ""

	for _, c := range s {
		isUpper := c >= 'A' && c <= 'Z'
		isLower := c >= 'a' && c <= 'z'

		if start && isLower {
			result += string(c - 32)
		} else if !start && isUpper {
			result += string(c + 32)
		} else {
			result += string(c)
		}
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') && (c < '0' || c > '9') {
			start = true
		} else {
			start = false
		}
	}

	return result
}
