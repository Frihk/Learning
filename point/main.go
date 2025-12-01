// package main

// // import "github.com/01-edu/z01"

// // type ptr struct {
// // 	x int
// // 	y int
// // }

// // func setPoint(pointer *ptr, newpointx int, newpointy int) {
// // 	pointer.x = newpointx
// // 	pointer.y = newpointy
// // }

// // func main() {
// // 	points := ptr{}
// // 	setPoint(&points, 42, 21)

// // }

// import "github.com/01-edu/z01"

// type point struct {
// 	x int
// 	y int
// }

// func setPoint(ptr *point) {
// 	ptr.x = 42
// 	ptr.y = 21
// }

// func print(s string) {
// 	for i := 0; i < len(s); i++ {
// 		z01.PrintRune(rune(s[i]))
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	p := &point{}
// 	setPoint(p)

// 	print("x = 42, y = 21") // works, but values are hardcoded
// }

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

// printString prints a string using z01.PrintRune
func printString(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

// printNumber prints a non-negative integer using z01.PrintRune
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Collect digits in reverse
	digits := []rune{}
	for n > 0 {
		d := n % 10
		digits = append(digits, rune('0'+d))
		n /= 10
	}

	// Print digits in correct order
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
