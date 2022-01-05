package missingnumber

func missingnumber(nums []int) int {
	current_sum := 0
	n := len(nums) + 1

	for _, val := range nums {
		current_sum += val
	}

	// trick is to use the arithmetic series formula
	actual_sum := (n * (n + 1)) / 2

	return Abs(actual_sum - current_sum)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
