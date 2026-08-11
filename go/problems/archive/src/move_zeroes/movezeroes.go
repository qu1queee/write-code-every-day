package movezeroes

func moveZeroes(nums []int) {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[read], nums[write] = nums[write], nums[read]
			write++
		}
	}
}
