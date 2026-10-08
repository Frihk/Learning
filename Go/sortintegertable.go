package piscine

func SortIntegerTable(table []int) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			temp := table[i]
			table[i] = table[i-1]
			table[i-1] = temp
			i = 1
		} else {
			i++
		}
	}
}
