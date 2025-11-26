package piscine

func ToUpper(s string) string {
	r := []rune(s)
	u := []rune{}
	for _, c := range r {
		if c >= 'a' && c <= 'z' {
			u = append(u, c-32)
		} else {
			u = append(u, c)
		}
	}
	return string(u)
}
