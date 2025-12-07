package piscine

import "github.com/01-edu/z01"

// first fuction to place the queen and back track if not safe //placer
// function to check for safety //checker
// function main to display the queens coordinates// main
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Checker(pos []int, row int, col int) bool {
	// loop through previous columns
	// check same row
	// check diagonal
	// return true if no conflict
	for i := 0; i < col; i++ {
		if pos[i] == row {
			return false
		}
		if abs(pos[i]-row) == abs(i-col) {
			return false
		}
	}
	return true
}

func Placer(positions []int, col int) {
	// loop through rows 1 to 8
	// check if row is safe
	// if safe, place queen and recurse to next column
	// if col == 8, print the solution
	for i := 1; i <= 8; i++ {
		if Checker(positions, i, col) {
			positions[col] = i
			if col == 7 {
				for j := 0; j < len(positions); j++ {
					z01.PrintRune(rune(positions[j] + '0'))
				}
				z01.PrintRune('\n')
			}
			Placer(positions, col+1)
			positions[col] = 0
		}
	}
}

func Eightofqueens() {
	var positions [8]int
	col := 0
	Placer(positions[:], col)
}
