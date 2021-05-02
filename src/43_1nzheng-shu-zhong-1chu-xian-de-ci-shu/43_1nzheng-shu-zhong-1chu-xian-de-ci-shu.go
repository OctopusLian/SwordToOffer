func countDigitOne(n int) int {
	// 这里 digitNum 为 digit 所在位 对应的可能性
	// 其实位为个位 因此 digitNum = 1  1乘以任何数 = 任何数
	digitNum, sum := 1, 0
	high, cur, low := n/10, n%10, 0
	for high != 0 || cur != 0 {
		// 固定位 计算数量
		if cur < 1 {
			sum += high * digitNum
		} else if cur == 1 {
			sum += high*digitNum + low + 1
		} else {
			sum += (high + 1) * digitNum
		}
		// 换下一位 更新高低位 及digit数量级
		low = low + cur*digitNum
		high, cur = high/10, high%10
		digitNum = digitNum * 10
	}
	return sum
}