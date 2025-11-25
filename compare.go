package piscine

func Compare(a, b string) int {
    c := len(a)
    d := len(b)

    // optional: handle empty strings safely
    if c == 0 || d == 0 {
        return 0
    }

    firstEqual := a[0] == b[0]
    lastEqual := a[c-1] == b[d-1]

    switch {
    case firstEqual && lastEqual:
        return 0
    case firstEqual && !lastEqual:
        return 1
    case !firstEqual && lastEqual:
        return -1
    case !firstEqual && !lastEqual:
        if a[0] > b[0] { // optional: decide which way to return
            return 1
        }
        return -1
    }

    return 0
}