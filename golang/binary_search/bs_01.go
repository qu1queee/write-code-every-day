package main

import "fmt"

func main() {
	fmt.Println(binary_search([]int{1, 3, 5, 7, 9}, 9))
	fmt.Println(binary_search([]int{1, 3, 5, 7, 9}, -1))
}

// We expect the array to be sorted
func binary_search(myarray []int, searchID int) int {
	low := 0
	high := len(myarray) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := myarray[mid]

		if guess == searchID {
			return mid
		}
		if guess > searchID {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}
