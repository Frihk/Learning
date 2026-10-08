package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, c := range args {
		if string(c) == "01" ||
			string(c) == "galaxy" ||
			string(c) == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
