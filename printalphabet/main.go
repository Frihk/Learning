package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// alphabets := 97
	for i := 97; i < 123; i++ {
		// alphabets := rune(i)
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
