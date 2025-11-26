package piscine


func Capitalize(s string) string {
    start := true
    result := ""

    for _, c := range s {
        isUpper := c >= 'A' && c <= 'Z'
        isLower := c >= 'a' && c <= 'z'

        if start && isLower {
            // Capitalize first letter of the word if it's lowercase
            result += string(c - 32)
        } else if !start && isUpper {
            // Lowercase all other letters
            result += string(c + 32)
        } else {
            result += string(c)
        }

        // Update start: only after non-letter AND non-digit
        if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') && (c < '0' || c > '9') {
            start = true
        } else {
            start = false
        }
    }

    return result
}