package piscine

func ActiveBits(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}
