package piscine

func StrRev(s string) string {
	norm := []rune(s)
	rev := []rune{}
	for i := len(norm) - 1; i >= 0; i-- {
		rev = append(rev, norm[i])
	}
	return string(rev)
}
