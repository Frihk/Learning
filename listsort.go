package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	var end *NodeI = nil
	for end != l.Next {
		current := l
		for current.Next != end {
			if current.Data > current.Next.Data {
				temp := current.Data
				current.Data = current.Next.Data
				current.Next.Data = temp
			}
			current = current.Next
		}
		end = current
	}
	return l
}
