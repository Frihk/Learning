package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	result := ""
	count := 0
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		result += string(str[i])
		count++
		if count == 5 {
			result += " "
			count = 0
			i++
		}
	}
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}
