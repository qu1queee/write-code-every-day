package lotz

import (
	"reflect"
	"testing"
)

func createBinaryTree(nums []int) *TreeNode {
	return createTree(nums, 0)
}

func createTree(nums []int, index int) *TreeNode {
	if index >= len(nums) {
		return nil
	}

	tree := TreeNode{Val: nums[index]}
	tree.Left = createTree(nums, 2*index+1)
	tree.Right = createTree(nums, 2*index+2)
	return &tree
}

func TestPreorderTraversal(t *testing.T) {
	testCases := []*TreeNode{
		createBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8}),
		createBinaryTree([]int{1, 2, 3, 4, 5}),
	}
	expected := [][][]int{
		{{1}, {3, 2}, {4, 5, 6, 7}, {8}},
		{{1}, {3, 2}, {4, 5}},
	}
	testFuncs := []func(node *TreeNode) [][]int{
		zigzagLevelOrder,
	}

	for _, testFunc := range testFuncs {
		for index, root := range testCases {
			if res := testFunc(root); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("got %v, expected %v", res, expected[index])
			}
		}
	}
}
