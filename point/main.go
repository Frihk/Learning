package main

// import "github.com/01-edu/z01"

// type ptr struct {
// 	x int
// 	y int
// }

// func setPoint(pointer *ptr, newpointx int, newpointy int) {
// 	pointer.x = newpointx
// 	pointer.y = newpointy
// }

// func main() {
// 	points := ptr{}
// 	setPoint(&points, 42, 21)

// }

import "github.com/01-edu/z01"

type ptr struct {
	x int
	y int
}

func setPoint(pointer *ptr, newpointx int, newpointy int) {
	pointer.x = newpointx
	pointer.y = newpointy
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func main() {
	points := ptr{}
	setPoint(&points, 42, 21)

	printStr("x = ")
	printNbr(points.x)
	z01.PrintRune('\n')

	printStr("y = ")
	printNbr(points.y)
	z01.PrintRune('\n')
}
