package lot

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	var result [][]int

	queue := []*TreeNode{root}

	for level := 0; len(queue) > 0; level++ {
		result = append(result, []int{})
		for elements := len(queue); elements > 0; elements-- {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			result[level] = append(result[level], queue[0].Val)
			queue = queue[1:]
		}
	}
	return result
}
