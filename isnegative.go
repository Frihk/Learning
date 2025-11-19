package main

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	mytruth := 'T'
	myfalse := 'F'
	if nb < 0 {
		k := rune(mytruth)
		z01.PrintRune(k)
		z01.PrintRune('\n')

	} else {
		k := rune(myfalse)
		z01.PrintRune(k)
		z01.PrintRune('\n')
	}
}
