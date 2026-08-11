package symmetrictreerecursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	return validateMirror(root, root)
}

func validateMirror(a *TreeNode, b *TreeNode) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return (a.Val == b.Val) && validateMirror(a.Right, b.Left) && validateMirror(a.Left, b.Right)
}
