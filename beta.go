package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// alphabets := 97
	for i := 97; i < 123; i++ {
		alphabets := string(i)
		z01.Print(alphabets)
	}
}


package main

import (
	"fmt"
)

func main() {
	// alphabets := 97
	for i := 97; i < 123; i++ {
		// alphabets += i
		fmt.Print(string(i))
	}
}
