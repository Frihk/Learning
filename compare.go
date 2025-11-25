package piscine

func Compare(a, b string) int {
  c := len(a)
  d := len(b)
  if a[c-1] != b[d-1] {
    return 1
  }
  if a[0] != b[0] {
    return -1
  }
  if a[0] == b[0] && a[c-1] == b[d-1] {
    return 0
  }
  return 0
}