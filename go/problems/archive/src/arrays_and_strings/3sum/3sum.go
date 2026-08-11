package threesum

// From Leetcode hhttps://leetcode.com/problems/3sum/

/*
Some personal notes:
- Tagged as Medium

Problem Complexity Analysis:
- Time Complexity: O(N^2)
- Space Complexity: O(N)

Additional Nodes:
- Trick is to use a hash similar to a 2sum, sort the list
  and avoid duplicates.
*/

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := new([][]int)

	sort.Ints(nums)

	for i := range nums {
		// following condition avoids duplicates,
		// it assumes the list is already sorted.
		if i == 0 || nums[i-1] != nums[i] {
			fmt.Println(i)
			twoSum(nums, i, result)
		}
	}
	return *result
}

func twoSum(nums []int, i int, result *[][]int) []int {
	mymap := make(map[int]int)
	for j := i + 1; j < len(nums); j++ {
		complement := -nums[i] - nums[j]
		if _, ok := mymap[complement]; ok {
			*result = append(*result, []int{nums[i], nums[j], complement})
			// following loop ensures that we avoid generation of duplicates
			for j+1 < len(nums) && nums[j] == nums[j+1] {
				j++
			}

		}
		mymap[nums[j]] = j
	}
	return nil
}
