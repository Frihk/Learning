package piscine

func Capitalize(s string) string {
	start := true
	result := ""
	for _, c := range s {
		if start && c >= 'a' && c <= 'z' {
			result += string(c - 32)
			start = false
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
