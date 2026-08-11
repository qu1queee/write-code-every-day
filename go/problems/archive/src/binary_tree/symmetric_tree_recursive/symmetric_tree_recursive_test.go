package symmetrictreerecursive

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
		createBinaryTree([]int{1}),
		createBinaryTree([]int{1, 2, 2}),
		createBinaryTree([]int{1, 1, 2}),
	}
	expected := []bool{
		true,
		true,
		false,
	}
	testFuncs := []func(root *TreeNode) bool{
		isSymmetric,
	}

	for _, testFunc := range testFuncs {
		for index, root := range testCases {
			if res := testFunc(root); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("got %v, expected %v", res, expected[index])
			}
		}
	}
}
