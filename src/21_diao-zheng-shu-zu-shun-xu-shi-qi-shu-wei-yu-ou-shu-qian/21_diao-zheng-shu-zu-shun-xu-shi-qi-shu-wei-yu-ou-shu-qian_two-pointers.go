func exchange(nums []int) []int {
	left, right := 0, len(nums)-1 //左右指针
	for left < right {            //结束条件
		for left < end && nums[left]%2 == 1 {
			// 正向遍历nums直到nums[left]为偶数，或者left>end
			left++
		}
		for left < right && nums[right]%2 == 0 {
			// 逆向遍历nums直到nums[right]]为奇数，或者left>end
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}

	return nums
}