package diameteroftree

// From Leetcode https://leetcode.com/problems/diameter-of-binary-tree/

/*
Some personal notes:
- Tagged as Easy

Problem Complexity Analysis:
- Time Complexity: O(N)
- Space Complexity: O(N)

Additional Nodes:
- Trick is to understand that the longest path will
  move from a leaf node to another leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	diameter := new(int)

	longestPath(root, diameter)

	return *diameter
}

func longestPath(root *TreeNode, diameter *int) int {

	if root == nil {
		return 0
	}

	var leftPath = longestPath(root.Left, diameter)

	var rightPath = longestPath(root.Right, diameter)

	if (leftPath + rightPath) > *diameter {
		*diameter = leftPath + rightPath
	}

	if leftPath > rightPath {
		return leftPath + 1
	} else {
		return rightPath + 1
	}

}
