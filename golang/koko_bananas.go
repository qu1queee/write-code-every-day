package main

import "fmt"

func main() {
	fmt.Println("foobar")
}

func minEatingSpeed(piles []int, hours int) int {

	var left = 1
	// maximum value of all piles
	var right = getMax(piles) + 1
	// can I finish bananas at X speed, in Y hours?
	// Let´s use binary search
	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, mid, hours) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func getMax(piles []int) int {
	return 0
}

func canFinish(piles []int, speed int, hours int) bool {
	return false
}
