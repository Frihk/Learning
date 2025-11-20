package main

import (
	"fmt"
)

func main() {
	n := 1
	j := &n
	fmt.Println(*j)
}
