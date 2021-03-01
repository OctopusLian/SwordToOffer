func findRepeatNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		tmp := nums[i]
		nums[i] = nums[tmp]
		nums[tmp] = tmp
	}

	return -1
}