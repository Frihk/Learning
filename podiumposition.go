package piscine

func PodiumPosition(podium [][]string) [][]string {
	slice := [][]string{}
	slice = append(slice, podium...)
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i][0] > slice[j][0] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}
