package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	s := make([]int, (max - min))
	for i := min; i < max; i++ {
		s[i-min] = i
	}
	return s
}
