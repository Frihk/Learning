package main

import (
	"os"
)

func write(s string) {
	os.Stdout.Write([]byte(s))
}

func writeln(s string) {
	os.Stdout.Write([]byte(s + "\n"))
}

func main() {
	arg := os.Args[1:]

	if len(arg) != 3 || len(arg[0]) > 10 || len(arg[2]) > 10 {
		return
	}
	num1 := 0
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
	num2 := 0
	sign = 1 // reset sign

	for i, c := range arg[2] {
		if i == 0 && c == '-' {
			sign = -1
			continue
		}
		if c < '0' || c > '9' {
			return
		}
		num2 = num2*10 + int(c-'0')
	}
	num2 *= sign
	op := arg[1]

	switch op {
	case "+":
		writeln(itoa(num1 + num2))

	case "-":
		writeln(itoa(num1 - num2))

	case "*":
		writeln(itoa(num1 * num2))

	case "%":
		if num2 == 0 {
			writeln("No modulo by 0")
		} else {
			writeln(itoa(num1 % num2))
		}

	case "/":
		if num2 == 0 {
			writeln("No division by 0")
		} else {
			writeln(itoa(num1 / num2))
		}

	default:
		return
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	buf := []byte{}
	for n > 0 {
		buf = append([]byte{byte('0' + (n % 10))}, buf...)
		n /= 10
	}

	return sign + string(buf)
}
