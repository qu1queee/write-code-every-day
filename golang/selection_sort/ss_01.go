package main

import "fmt"

func main() {
	fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
}

func findSmallest(list []int) int {
	smallest := list[0]
	smallest_index := 0
	for i := 1; i < len(list); i++ {
		if list[i] < smallest {
			smallest = list[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectionSort(list []int) []int {
	newList := []int{}
	for range list {
		smallest := findSmallest(list)
		newList = append(newList, list[smallest])
		list = append(list[:smallest], list[smallest+1:]...)
	}
	return newList
}
