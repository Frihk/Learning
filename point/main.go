package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// printString prints a string using only z01.PrintRune
func printString(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

// printNumber prints a non-negative integer using only z01.PrintRune
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune{}
	for n > 0 {
		digits = append(digits, rune('0'+(n%10)))
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func main() {
	p := &point{}
	setPoint(p)

	printString("x = ")
	printNumber(p.x)
	printString(", y = ")
	printNumber(p.y)
	z01.PrintRune('\n')
}
