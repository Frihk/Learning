package piscine

func Abort(a, b, c, d, e int) int {
	sor := []int{}
	sor = append(sor, a, b, c, d, e)
	for i := 0; i < len(sor); i++ {
		for j := i + 1; j < len(sor); j++ {
			if sor[i] > sor[j] {
				temp := sor[i]
				sor[i] = sor[j]
				sor[j] = temp
			}
		}
	}
	k := len(sor) / 2
	return sor[k]
}

// func main() {
// 	middle := Abort(2, 3, 8, 5, 7)
// 	fmt.Println(middle)
// }
