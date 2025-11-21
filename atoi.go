package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	for i := start; i < len(s); i++ {
		r := s[i]
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
		} else {
			return 0 // invalid character
		}
	}

	return result * sign
}
