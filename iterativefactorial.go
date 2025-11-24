package piscine

func IterativeFactorial(nb int) int {
	fact := 1
	if nb < 0 {
		return 0
	}
	if nb > 0 {
		for i := 1; i <= nb; i++ {
			fact *= i
		}
	}
	return fact
}
