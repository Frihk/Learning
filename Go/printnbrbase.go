package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	basel := len(base)

	// Base must have at least 2 characters
	if basel < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	// Validate base
	for i := 0; i < basel; i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < basel; j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	// Convert to uint64 safely
	var unbr uint64
	if nbr < 0 {
		z01.PrintRune('-')
		unbr = uint64(-(nbr + 1)) + 1 // safe conversion of MinInt
	} else {
		unbr = uint64(nbr)
	}

	// Write digits
	var digits []rune
	if unbr == 0 {
		digits = append(digits, rune(base[0]))
	}
	for unbr > 0 {
		rem := unbr % uint64(basel)
		digits = append(digits, rune(base[rem]))
		unbr /= uint64(basel)
	}

	// Print digits in reverse
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
