package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune('0')
	}
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	i := 1
	for i < len(digits) {
		if digits[i-1] > digits[i] {
			digits[i-1], digits[i] = digits[i], digits[i-1]
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}

	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
	z01.PrintRune('\n')
}
