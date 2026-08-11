package preorder_traversal

// From Leetcode https://leetcode.com/problems/binary-tree-preorder-traversal/

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

func preorderTraversal(root *TreeNode) []int {

	var stack []*TreeNode
	var output []int

	if root == nil {
		return output
	}

	stack = append(stack, root)

	for len(stack) > 0 {

		current := stack[0]
		stack = stack[1:]

		output = append(output, current.Val)

		if current.Right != nil {
			stack = append([]*TreeNode{current.Right}, stack...)
		}

		if current.Left != nil {
			stack = append([]*TreeNode{current.Left}, stack...)
		}

	}
	return output
}
