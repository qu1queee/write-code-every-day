package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 4, 5, 9, 23, 12, 45}, 45))
	fmt.Println(binarySearchLeftBorder([]int{2, 3, 5, 7}, 3))
	fmt.Println(binarySearchRightBorder([]int{2, 3, 6}, 6))

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

func binarySearchLeftBorder(nums []int, target int) int {
	var left int = 0
	var right int = len(nums) - 1

	// search interval is [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// search interval is [mid+1,right]
			left = mid + 1
		} else if nums[mid] > target {
			// search interval is [left, mid-1]
			right = mid - 1
		} else if nums[mid] == target {
			// shrink right border
			right = mid - 1
		}
	}
	// when target is greater than all elements, left will be
	// get out of bounds, equal to len(nums), therefore we need
	// to prevent it
	if left >= len(nums) || nums[left] != target {
		return -1
	}

	return left
}

func binarySearchRightBorder(nums []int, target int) int {
	var left int = 0
	var right int = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			// short to the right side
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// when target is smaller than all elements, right will
	// be reduced to -1 , therfore we need prevent going out
	// of bounds
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
