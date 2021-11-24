package sumclosest

// From Leetcode https://leetcode.com/problems/3sum-closest/

/*
Some personal notes:
- Tagged as Medium

Problem Complexity Analysis:
- Time Complexity: O(N^2)
- Space Complexity: O(logN) to O(N), dependes on the sorting algorithm

Additional Nodes:
- Trick is manage two pointers ( low and high )
*/

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	diff := math.MaxInt64
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		lo := i + 1
		hi := len(nums) - 1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if math.Abs(float64(target)-float64(sum)) < math.Abs(float64(diff)) {
				diff = target - sum
			}
			if sum < target {
				lo++
			} else {
				hi--
			}
			if diff == 0 {
				break
			}
		}
	}

	return target - diff
}
