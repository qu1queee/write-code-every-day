package inorder_traversal

// From Leetcode https://leetcode.com/problems/binary-tree-inorder-traversal/

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

func inorderTraversal(root *TreeNode) []int {

	var stack []*TreeNode
	var list []int

	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// pop last item in stack
		popIndex := len(stack) - 1
		current = stack[popIndex]
		stack = stack[0:popIndex]
		list = append(list, current.Val)
		current = current.Right

	}

	return list
}
