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
	arguments := os.Args

	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if i != 0 && (runes[j] >= '0' && runes[j] <= '9') {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if i != 0 && (runes[j] >= 'A' && runes[j] <= 'Z') {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if i != 0 && (runes[j] >= 'a' && runes[j] <= 'z') {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
}
