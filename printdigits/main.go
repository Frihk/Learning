package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// alphabets := 97
	for i := 48; i <= 57; i++ {
		// alphabets := rune(i)
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
