package piscine

func AtoiBase(s string, base string) int {
    baselen := len(base)
    if baselen < 2 {
        return 0
    }

    // Validate base: no '+' or '-', no duplicates
    for i := 0; i < baselen; i++ {
        if base[i] == '+' || base[i] == '-' {
            return 0
        }
        for j := i + 1; j < baselen; j++ {
            if base[i] == base[j] {
                return 0
            }
        }
    }

    // Convert s from base to integer
    result := 0

    for _, c := range s {
        // Find index of c in base
        idx := -1
        for i := 0; i < baselen; i++ {
            if rune(base[i]) == c {
                idx = i
                break
            }
        }
        if idx == -1 { // character not found in base
            return 0
        }

        result = result*baselen + idx
    }

    return result
}
