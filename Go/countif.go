package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, c := range tab {
		if f(c) {
			counter += 1
		}
	}
	return counter
}

// func main() {
// 	a1 := []string{"This", "1", "is", "4", "you"}
// 	a2 := []string{"This", "is", "4", "you"}

// 	result1 := CountIf(IsNumeric, a1)
// 	result2 := CountIf(IsNumeric, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
