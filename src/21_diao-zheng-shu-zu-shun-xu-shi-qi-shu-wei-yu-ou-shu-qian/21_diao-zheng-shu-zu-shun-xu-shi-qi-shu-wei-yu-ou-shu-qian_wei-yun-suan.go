func exchange(nums []int) []int {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

//链接：https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/solution/shi-jian-fu-za-du-on-kong-jian-fu-za-du-o1-zhi-sao/
