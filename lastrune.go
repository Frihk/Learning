package piscine

func LastRune(s string) rune {
  r := []rune(s)
  l := len(s) -1
  return r[l]
}
