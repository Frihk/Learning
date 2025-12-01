package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printError(err error) {
	for _, c := range "ERROR: " + err.Error() {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
	os.Exit(1)
}

func printBytes(b []byte) {
	for _, c := range b {
		z01.PrintRune(rune(c))
	}
}

func main() {
	args := os.Args

	// No arguments → read from stdin
	if len(args) == 1 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			printError(err)
		}
		printBytes(data)
		return
	}

	// With arguments → cat each file
	for i := 1; i < len(args); i++ {
		f, err := os.Open(args[i])
		if err != nil {
			printError(err)
		}

		data, err := io.ReadAll(f)
		f.Close() // close immediately, not via defer
		if err != nil {
			printError(err)
		}

		printBytes(data)
	}
}
