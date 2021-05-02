func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	window, res := []int{}, []int{}

	for i, x := range nums {
		// 形成窗口，且在窗口范围之外了，移除
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}

		// 开始维护队列操作，队列末端 如果 小于 待进入元素，则直接出队
		for len(window) != 0 && nums[window[len(window)-1]] <= x {
			window = window[:len(window)-1]
		}

		window = append(window, i)

		// 窗口已形成之后，可以开始输出结果集了
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}

	}
	return res
}

//链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/solution/goshuang-duan-dui-lie-shi-xian-cun-chu-s-8pnc/
