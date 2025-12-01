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

// import "github.com/01-edu/z01"

// type ptr struct {
// 	x int
// 	y int
// }

// func setPoint(ptr *ptr) {
// 	ptr.x = 42
// 	ptr.y = 21
// }

// func main() {
// 	point := ptr{}
// 	setPoint(&point)

// 	z01.PrintRune('x')
// 	z01.PrintRune(' ')
// 	z01.PrintRune('=')
// 	z01.PrintRune(' ')
// 	z01.PrintRune('4')
// 	z01.PrintRune('2')
// 	z01.PrintRune('\n')
// 	z01.PrintRune('y')
// 	z01.PrintRune(' ')
// 	z01.PrintRune('=')
// 	z01.PrintRune(' ')
// 	z01.PrintRune('2')
// 	z01.PrintRune('1')
// 	z01.PrintRune('\n')
// }

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('4')
	z01.PrintRune('2')
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('2')
	z01.PrintRune('1')
	z01.PrintRune('\n')
}
