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

func SortTable(args []string) {
	for i := 0; i < len(args)-1; i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	SortTable(args)
	for _, arg := range args {
		for _, b := range arg {
			z01.PrintRune(rune(b))
		}
		z01.PrintRune('\n')
	}
}
