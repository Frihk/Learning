package piscine

func Any(f func(string) bool, a []string) bool {
	// str := ""
	for _, c := range a {
		if f(c) {
			return false
		}
	}
	return true
}

// func main() {
// 	a1 := []string{"Hello", "how", "are", "you"}
// 	a2 := []string{"This", "is", "4", "you"}

// 	result1 := piscine.Any(IsNumeric, a1)
// 	result2 := piscine.Any(IsNumeric, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
