package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data > n2.Data {
		n1, n2 = n2, n1
	}
	head := n1
	for n1.Next != nil && n2 != nil {
		if n1.Next.Data > n2.Data {
			temp := n2.Next
			n2.Next = n1.Next
			n1.Next = n2
			n2 = temp
		}
		n1 = n1.Next
	}
	if n2 != nil {
		n1.Next = n2
	}

	return head
}
