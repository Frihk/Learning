package piscine

import "fmt"

func PrintWordsTables(a []string) {
	results := ""
	for i, c := range a {
		results += c
		if i < len(a)-1 {
			results += "\n"
		}
	}
	fmt.Print(results)
}
