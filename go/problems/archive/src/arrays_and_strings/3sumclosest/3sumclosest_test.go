package sumclosest

import (
	"reflect"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	testCases := [][]int{
		{-1, 2, 1, -4},
		{0, 0, 0},
	}
	targetForCase := []int{
		1,
		1,
	}
	expected := []int{
		2,
		0,
	}
	testFuncs := []func(nums []int, target int) int{
		threeSumClosest,
	}

	for _, testFunc := range testFuncs {
		for index, root := range testCases {
			if res := testFunc(root, targetForCase[index]); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("got %v, expected %v", res, expected[index])
			}
		}
	}
}
