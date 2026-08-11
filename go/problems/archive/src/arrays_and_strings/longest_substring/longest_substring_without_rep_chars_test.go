package lsworc

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	testCases := []string{
		"pwwkew",
		"bbbbb",
		"abcabcbb",
	}
	expected := []int{
		3,
		1,
		3,
	}
	testFuncs := []func(s string) int{
		lengthOfLongestSubstring,
	}

	for _, testFunc := range testFuncs {
		for index, root := range testCases {
			if res := testFunc(root); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("got %v, expected %v", res, expected[index])
			}
		}
	}
}
