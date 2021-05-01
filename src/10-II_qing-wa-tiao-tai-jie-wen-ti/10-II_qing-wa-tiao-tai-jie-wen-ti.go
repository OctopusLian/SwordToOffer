//与斐波那契数列解法一致
func numWays(n int) int {
	a := 1
	b := 1
	var sum int
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}

	return a
}