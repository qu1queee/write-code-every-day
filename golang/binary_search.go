package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 4, 5, 9, 23, 12, 45}, 45))
}

func binarySearch(nums []int, target int) int {
	var left int = 0
	var right int = len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
