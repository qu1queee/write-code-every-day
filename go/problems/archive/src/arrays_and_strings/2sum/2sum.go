package twosum

func twosum(list []int, target int) bool {

	myMap := make(map[int]int)

	for _, val := range list {
		if _, exists := myMap[target-val]; exists {
			return true
		} else {
			myMap[val] = 0
		}

	}
	return false
}
