// 动态规划
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1 //当绳子长为0， 1， 2的情况
	var temp1, temp2 int
	for i := 3; i < n+1; i++ {
		//将绳子分成i-1，1，两段，可以看出dp[i]必然大于等于dp[i-1]
		dp[i] = dp[i-1]
		for j := 1; j <= i/2; j++ {
			//将绳子分为两部分，两部分自由组合分段然后与不分段比较，选择较大的两部分组合相乘
			temp1 = max(dp[j], j)
			temp2 = max(dp[i-j], i-j)
			dp[i] = max(dp[i], temp1*temp2)
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}