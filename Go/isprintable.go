package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	// pri := []rune{}
	for _, c := range r {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}
