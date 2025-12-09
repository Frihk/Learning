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
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data > root.Data {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	} else if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return root
}
func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Data >= root.Data || !BTreeIsBinary(root.Left) {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Data <= root.Data || !BTreeIsBinary(root.Right) {
			return false
		}
	}
	return true
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