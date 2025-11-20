package main

// import (
// 	"fmt"
// 	"slices"
// )

// 	func append_myslice(num ...int) {
// 		myslice := []int{}
// 		myslice = append(myslice, num...)
// 		slices.Sort(myslice)
// 		for i := 0; i < len(myslice); i++ {
// 			fmt.Print(myslice[i], "\n")
// 		}
// 	}
func main() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				z01.PrintRune('')
			}
		}
	}
}
import (
	"fmt"
	"slices"
)



func main() {
	myslice := []int{} 
	num = 
	myslice = append(myslice, num...)
	slices.Sort(myslice)
	for i := 0; i < len(myslice); i++ {
		new_slice := rune(myslice[i])
		fmt.Print(new_slice, " ")
}
