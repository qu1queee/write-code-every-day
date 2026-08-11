package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := [][]int{
		{3, 2, 4},
		{2, 7, 11, 15},
	}
	targetForCase := []int{
		6,
		9,
	}

	expected := []bool{
		true,
		true,
	}

	testFuncs := []func(s []int, a int) bool{
		twosum,
	}

	for _, testFunc := range testFuncs {
		for index, root := range testCases {
			if res := testFunc(root, targetForCase[index]); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("got %v, expected %v", res, expected[index])
			}
		}
	}
}
