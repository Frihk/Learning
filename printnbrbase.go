package piscine

import (
	"github.com/01-edu/z01"
	// "github.com/01-edu/z01.PrintRune"
)

func PrintNbrBase(nbr int, base string) {
	basel := len(base)
	digits := []rune{}
	if basel < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < basel; i++ {
		for j := 0; j < basel; j++ {
			if i != j && base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
		if base[i] == '-' || base[i] == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	for nbr > 0 {
		rem := nbr % basel
		digits = (append(digits, rune(base[rem])))
		nbr = nbr / len(base)
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}

}