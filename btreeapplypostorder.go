package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	new := []*TreeNode{root}
	var lastVisited *TreeNode

	for len(new) > 0 {
		current := new[len(stack)-1]

		if lastVisited == nil || lastVisited.Left == current || lastVisited.Right == current {
			if current.Left != nil {
				new = append(new, current.Left)
			} else if current.Right != nil {
				new = append(new, current.Right)
			}
		} else if current.Left == lastVisited {
			if current.Right != nil {
				new = append(new, current.Right)
			}
		} else {
			f(current.Data)
			new = new[:len(new)-1]
		}

		lastVisited = current
	}
}
