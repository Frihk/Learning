package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		result := 1
		for i := 1; i < nb+1; i++ {
			result *= i
			if result > 2147483647 {
				result = 0
				break
			}
		}
		return result
	}
}

// func IterativeFactorial(nb int) int {
// 	// fact := 1
// 	if nb < 0 {
// 		return 0
// 	}
// 	if nb == 0 ||
// 		nb == 1 {
// 		return 1
// 	}
// 	return nb * IterativeFactorial(nb-1)
// }
