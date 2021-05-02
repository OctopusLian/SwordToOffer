func majorityElement(nums []int) int {
	res, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			res = num
			cnt += 1
		} else {
			if num == res {
				cnt += 1
			} else {
				cnt -= 1
			}
		}
	}
	return res
}