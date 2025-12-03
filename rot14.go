package piscine

func Rot14(s string) string {
  str := ""
  for _, c := range s{
    if c >= 'a'&& c <= 'z' {
      str = str + string((c-'a'+14)%26 + 'a')
    }else if c >= 'A'&& c <= 'Z' {
      str = str + string((c-'A'+14)%26 + 'A')
    }else {
      str = str + string(c)
    }
  }
  return str
}