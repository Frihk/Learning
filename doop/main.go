package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 || len(arg[0]) > 10 || len(arg[2]) > 10 {
		return
	}
	num1 := 0
	num2 := 0
	sign := 1
	for i, c := range arg[0] {
		if i == 0 && c == '-' {
			sign = -1
			continue
		}
		if c < '0' || c > '9' {
			return
		}
		num1 = num1*10 + int(c-'0')
	}
	num1 *= sign
	for i, d := range arg[2] {
		// handle negative sign
		if i == 0 && d == '-' {
			sign = -1
			continue
		}
		if d < '0' || d > '9' {
			return
		}
		num2 = num2*10 + int(d-'0')
	}
	num2 *= sign
	if arg[1] == "+" {
		fmt.Println(num1 + num2)
	}
	if arg[1] == "-" {
		fmt.Println(num1 - num2)
	}
	if arg[1] == "*" {
		fmt.Println(num1 * num2)
	}
	if arg[1] == "%" {
		if num2 == 0 {
			fmt.Print("No modulo by 0")
			fmt.Println()
		} else {
			fmt.Println(num1 % num2)
		}
	}
	if arg[1] == "/" {
		if num2 == 0 {
			fmt.Print("No division by 0")
		} else {
			fmt.Println(num1 / num2)
		}
	} else {
		return
	}
}
