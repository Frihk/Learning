package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	current := l
	for current != nil {
		if count == pos {
			return current
		}
		current = current.Next
		count++

	}
	return nil
}
