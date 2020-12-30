package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
}

func unitConverter(current int) int {
	// 1 foot -> 12 inches
	// 1 yard -> 3 feet
	// 1 chain -> 22 yards -> 66 feet
	return 0
}

// The United States uses the imperial system of weights and measures, which means that there are many different, seemingly arbitrary units to measure distance.
// There are 12 inches in a foot, 3 feet in a yard, 22 yards in a chain, and so on.
// Create a data structure that can efficiently convert a certain quantity of one unit
// to the correct amount of any other unit. You should also allow for additional units
// to be added to the system.
