package piscine

func Compact(ptr *[]string) int {
	part := []string{}
	for _, i := range *ptr {
		if i != "" {
			part = append(part, i)
		}
	}
	*ptr = part
	return len(part)
}
