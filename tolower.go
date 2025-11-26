package piscine

func ToLower(s string) string {
	r := []rune(s)
	u := []rune{}
	for _, c := range r {
		if c >= 'A' && c <= 'Z' {
			u = append(u, c+32)
		} else {
			u = append(u, c)
		}
	}
	return string(u)
}
