package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j < '9'; j++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			if i != '7' || j != '8' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

			// z01.PrintRune('\n')
		}
		// z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
