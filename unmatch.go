package piscine

func Unmatch(a []int) int {
	result := 0
	for _, c := range a {
		result ^= c
	}
	if result == 0 {
		return -1
	}
	return result
}
