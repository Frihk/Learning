package piscine


type TreeNode struct {
    Left, Right, Parent *TreeNode
    Data                string
}

func BTreeMin(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    current := root
    for current.Left != nil {
        current = current.Left
    }
    return current
}

func BTreeMax(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    current := root
    for current.Right != nil {
        current = current.Right
    }
    return current
}


func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		if root.Right != nil {
			min := BTreeMin(root.Right)
			root.Data = min.Data
			node = min
		} else if root.Left != nil {
			max := BTreeMax(root.Left)
			root.Data = max.Data
			node = max
		} else {
			return nil
		}
	}
	root.Left = BTreeDeleteNode(root.Left, node)
	root.Right = BTreeDeleteNode(root.Right, node)
	return root
}