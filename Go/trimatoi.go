package piscine

func TrimAtoi(s string) int {
	sign := 1
	num := 0
	firstDigitIndex := -1
	signIndex := -1

	// Find the index of the first digit and the first sign
	for i, c := range s {
		if c >= '0' && c <= '9' && firstDigitIndex == -1 {
			firstDigitIndex = i
		}
		if (c == '-' || c == '+') && signIndex == -1 {
			signIndex = i
		}
	}

	// Determine the sign: if '-' comes before the first digit
	if signIndex != -1 && firstDigitIndex != -1 && signIndex < firstDigitIndex {
		if s[signIndex] == '-' {
			sign = -1
		}
	}

	// Collect all digits
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
	}

	return num * sign
}
