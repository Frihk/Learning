package piscine

func DescendAppendRange(max, min int) []int {
	s := []int{}
	if max <= min {
		return nil
	}
	for i := max; i > min; i-- {
		s = append(s, i)
	}
	return s
}
