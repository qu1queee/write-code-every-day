package tdr

import "math"

// From Leetcode https://leetcode.com/problems/maximum-depth-of-binary-tree/

/*
Some personal notes:
- Tagged as Easy
- Code applies a recursive approach

Problem Complexity Analysis:
- Time Complexity: O(N)
- Space Complexity: log(N) (best scenario, tree is balanced)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	var answer = new(int)

	recursiveDepthEstimator(root, 1, answer)

	return *answer

}

func recursiveDepthEstimator(root *TreeNode, depth int, answer *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil { // leaf node
		*answer = int(math.Max(float64(*answer), float64(depth)))
	}

	recursiveDepthEstimator(root.Left, depth+1, answer)
	recursiveDepthEstimator(root.Right, depth+1, answer)
}
