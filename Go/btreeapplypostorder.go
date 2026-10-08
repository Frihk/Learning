package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	new := []*TreeNode{root}
	var lastnum *TreeNode

	for len(new) > 0 {
		current := new[len(new)-1]

		if lastnum == nil ||
			lastnum.Left == current ||
			lastnum.Right == current {
			if current.Left != nil {
				new = append(new, current.Left)
			} else if current.Right != nil {
				new = append(new, current.Right)
			}
		} else if current.Left == lastnum {
			if current.Right != nil {
				new = append(new, current.Right)
			}
		} else {
			f(current.Data)
			new = new[:len(new)-1]
		}

		lastnum = current
	}
}

// func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}
// 	BTreeApplyPostorder(root.Left, f)
// 	BTreeApplyPostorder(root.Right, f)
// 	f(root.Data)
// }
