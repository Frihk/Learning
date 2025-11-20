package piscine

import (
	"strconv"

	"github.com/01-edu/z01"
)

func PrintNbrStr(n int) {
	str := strconv.Itoa(n) // convert int to string
	for _, c := range str {
		z01.PrintRune(c)
	}
}
