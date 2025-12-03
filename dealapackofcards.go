package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	player1 := deck[:3]
	player2 := deck[3:6]
	player3 := deck[6:9]
	player4 := deck[9:12]

	fmt.Printf("player 1: ")
	for _, i := range player1 {
		fmt.Printf("%d, ", i)
	}
	z01.PrintRune('\n')

	fmt.Printf("player 2: ")
	for _, j := range player2 {
		fmt.Printf("%d, ", j)
	}
	z01.PrintRune('\n')

	fmt.Printf("player 3: ")
	for _, k := range player3 {
		fmt.Printf("%d, ", k)
	}
	z01.PrintRune('\n')

	fmt.Printf("player 4: ")
	for _, l := range player4 {
		fmt.Printf("%d, ", l)
	}
	z01.PrintRune('\n')
}
