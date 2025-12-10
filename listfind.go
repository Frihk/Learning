package piscine

func CompInt(a, b interface{}) bool {
	ai, ok1 := a.(int)
	bi, ok2 := b.(int)
	if !ok1 || !ok2 {
		return false
	}
	return ai == bi
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l == nil || l.Head == nil {
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
