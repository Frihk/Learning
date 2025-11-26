package piscine


func Capitalize(s string) string {
    start := true
    result := ""

    for _, c := range s {
        isUpper := c >= 'A' && c <= 'Z'
        isLower := c >= 'a' && c <= 'z'

        if start {
            if isLower {
                result += string(c - 32) 
            } else {
                result += string(c) 
            }
            start = false
        } else {
            if isUpper {
                result += string(c + 32) 
            } else {
                result += string(c)
            }
        }
        if !(isUpper || isLower) {
            start = true
        } else {
            start = false
        }
    }
    return result
}

