package piscine

func Compare(a, b string) int {
  c := len(a)
  d := len(b)
  if a[0] == b[0] && a[c-1] != b[d-1] {
    return 1
  }
  if a[0] != b[0] && a[c-1] == b[d-1]{
    return -1
  }
  if a[0] == b[0] && a[c-1] == b[d-1] {
    return 0
  }
   if a[0] != b[0] && a[c-1] != b[d-1] { // <-- added this case
    return -1
  }
  return 0
}