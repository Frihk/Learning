package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	BTreeApplyInorder(root.Left, f)
	BTreeApplyInorder(root.Right, f)
	f(root.Data)
}
