package validate_bst

// From Leetcode https://leetcode.com/problems/validate-binary-search-tree/

/*
Some personal notes:
- Tagged as Medium
- Code applies a recursive approach

Problem Complexity Analysis:
- Time Complexity: O(N) as we visit each node once
- Space Complexity: O(N)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	return isBST(root, nil, nil)
}

func isBST(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	return isBST(root.Left, min, root) && isBST(root.Right, root, max)

}
