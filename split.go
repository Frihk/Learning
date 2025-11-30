package piscine

func Split(s, sep string) []string {
    results := []string{}
    current := ""
    sepl := len(sep)
    for i := 0; i<len(s);{
        if i+sepl <= len(s) && s[i:i+sepl] == sep{
            results = append(results, current)
            current = ""
            i += sepl
        }else {
                current += string(s[i])
                i++
        }
    }
  results = append(results, current)
  return results
}