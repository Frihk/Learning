package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompInt(a, b interface{}) bool {
	ai, ok1 := a.(int)
	bi, ok2 := b.(int)
	if !ok1 || !ok2 {
		return false
	}
	return ai == bi
}

func CompStr(a, b interface{}) bool {
	s1, ok1 := a.(string)
	s2, ok2 := b.(string)
	if !ok1 || !ok2 {
		return false
	}
	return s1 == s2
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
