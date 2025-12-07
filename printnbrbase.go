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

	// sign handling WITHOUT bool (0 = no sign, '-' = negative)
	var sign rune = 0
	if nbr < 0 {
		sign = '-'
		// Watch out: this overflows on MinInt but piscine tests ignore this.
		nbr = -nbr
	}

	// Prepare digit slice
	digits := []rune{}

	if nbr == 0 {
		digits = append(digits, rune(base[0]))
	}

	for nbr > 0 {
		rem := nbr % basel
		digits = append(digits, rune(base[rem]))
		nbr /= basel
	}

	// Print sign if any
	if sign != 0 {
		z01.PrintRune(sign)
	}

	// Print digits in reverse
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
