package main

import (
	"os"

	"github.com/01-edu/z01"
)

/*func main() {
	args := os.Args[1:]
	i := 1
	for i < len(args) { // keep your loop
		if args[i-1] > args[i] { // compare safely
			// swap
			temp := args[i-1]
			args[i-1] = args[i]
			args[i] = temp

			if i > 1 { // gnome step back
				i--
			}
		} else {
			i++ // move forward
		}
	}
	for j, arg := range args {
		for _, c := range arg {
			z01.PrintRune(c)
		}
		if j < len(args)-1 {
			z01.PrintRune(' ') // space between arguments
		}
		z01.PrintRune('\n') // end with newline
	}
} */

func main() {
	a := os.Args[1:]
	for x, i := range a {
		for y, j := range a {
			if i < j {
				a[x], a[y] = a[y], a[x]
			}
		}
	}
	for i := 0; i < len(a); i++ {
		r := []rune(a[i])

		if r[0] < 'A' {
			for j := range a[i] {
				z01.PrintRune(r[j])
			}
			z01.PrintRune('\n')
		}

	}

	for i := 0; i < len(a); i++ {
		r := []rune(a[i])

		if r[0] >= '0' && r[0] <= '9' {
			for j := range a[i] {
				z01.PrintRune(r[j])
			}
			z01.PrintRune('\n')
		}

	}
	for i := 0; i < len(a); i++ {
		r := []rune(a[i])

		if r[0] >= 'A' && r[0] <= 'Z' {
			for j := range a[i] {
				z01.PrintRune(r[j])
			}
			z01.PrintRune('\n')
		}

	}
	for i := 0; i < len(a); i++ {
		r := []rune(a[i])

		if r[0] >= 'a' && r[0] <= 'z' {
			for j := range a[i] {
				z01.PrintRune(r[j])
			}
			z01.PrintRune('\n')
		}

	}
}
