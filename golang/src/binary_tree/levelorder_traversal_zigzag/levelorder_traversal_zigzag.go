package lotz

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for queueLevel := 0; len(queue) > 0; queueLevel++ {
		result = append(result, []int{})
		for items := len(queue) - 1; items >= 0; items-- {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if (queueLevel % 2) > 0 {
				result[queueLevel] = append([]int{queue[0].Val}, result[queueLevel]...)
			} else {
				result[queueLevel] = append(result[queueLevel], queue[0].Val)
			}

			queue = queue[1:]
		}

	}
	return result
}
