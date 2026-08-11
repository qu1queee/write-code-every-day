package postorder_traversal

// From Leetcode https://leetcode.com/problems/binary-tree-postorder-traversal/

/*
Some personal notes:
- Tagged as Easy
- Code does the iterative approach

Problem Complexity Analysis:
- Time Complexity: O(N)
- Space Complexity: O(N)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		if root != nil {
			result = append([]int{root.Val}, result...)
			stack = append(stack, root)
			root = root.Right
		} else {
			root = stack[len(stack)-1].Left
			stack = stack[:len(stack)-1]
		}
	}
	return result
}
