package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""

	for _, c := range s {
		if c == ' ' || c == '\n' || c == '\t' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}
