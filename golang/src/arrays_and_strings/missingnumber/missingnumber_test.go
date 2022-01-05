package missingnumber

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	testCases := [][]int{
		{3, 7, 1, 2, 8, 4, 5},
		{2, 3, 5, 6, 7, 9, 8},
	}
	expected := []int{
		6,
		4,
	}
	testFuncs := []func(s []int) int{
		missingnumber,
	}

	for _, testFunc := range testFuncs {
		for index, root := range testCases {
			if res := testFunc(root); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("got %v, expected %v", res, expected[index])
			}
		}
	}
}
