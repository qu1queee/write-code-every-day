package validate_bst

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
		createBinaryTree([]int{2, 1, 3}),
		createBinaryTree([]int{4, 6, 3}),
		createBinaryTree([]int{2}),
	}
	expected := []bool{
		true,
		false,
		true,
	}
	testFuncs := []func(root *TreeNode) bool{
		isValidBST,
	}

	for _, testFunc := range testFuncs {
		for index, root := range testCases {
			if res := testFunc(root); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("got %v, expected %v", res, expected[index])
			}
		}
	}
}
