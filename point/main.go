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

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	// detect --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, s := range args { // iterate actual arguments
		c := 0
		for _, r := range s { // iterate runes
			if r >= '0' && r <= '9' {
				c = c*10 + int(r-'0')
			}
		}

		if !(c >= 1 && c <= 26) {
			z01.PrintRune(' ')
		} else if upper {
			z01.PrintRune(rune('A' + (c - 1)))
		} else {
			z01.PrintRune(rune('a' + (c - 1)))
		}
	}
	z01.PrintRune('\n')
}
