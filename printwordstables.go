package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	results := ""
	for _, c := range a {
		results = results + c + "\n"
		// if i < len(a)-1 {
		// 	results += "\n"
		// }
	}
	for _, r := range results {
		z01.PrintRune(r)
	}
}
