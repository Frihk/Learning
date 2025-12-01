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

// printString prints a string using only z01.PrintRune
func Displaystr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

// printNumber prints a non-negative integer using only z01.PrintRune
func Diplaynum(n int) {
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

	Displaystr("x = ")
	Diplaynum(p.x)
	Displaystr(", y = ")
	Diplaynum(p.y)
	z01.PrintRune('\n')
}
