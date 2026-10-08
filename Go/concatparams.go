package piscine

func ConcatParams(args []string) string {
	results := ""
	for i, c := range args {
		results += c
		if i < len(args)-1 {
			results += "\n"
		}
	}
	return results
}
