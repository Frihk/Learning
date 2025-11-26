package piscine

import "fmt"

func PrintNbrInOrder(n int) {
	ord2 := fmt.Sprint(n)
	ord := []rune(ord2)

	i := 1
	for i < len(ord) {
		if ord[i-1] > ord[i] {
			temp := ord[i]
			ord[i] = ord[i-1]
			ord[i-1] = temp
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}
	sortedStr := string(ord)
	fmt.Println(sortedStr)
}
