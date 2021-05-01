//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

func myPow(x float64, n int) float64 {
	if n < 0 { // n为负数，进行转换
		x = 1 / x
		n = -n
	}
	pow := float64(1)
	for n != 0 {
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}