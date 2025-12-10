package piscine

func CompStr(a, b interface{}) bool {
	s1, ok1 := a.(string)
	s2, ok2 := b.(string)
	if !ok1 || !ok2 {
		return false
	}
	return s1 == s2
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l == nil {
		return nil
	}

	current := l.Head
	for current != nil {
		if comp(current.Data, ref) {
			return &current.Data
		}
		current = current.Next
	}

	return nil
}
