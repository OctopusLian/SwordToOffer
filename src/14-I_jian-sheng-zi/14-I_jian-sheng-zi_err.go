package main

import "math"

func cuttingRope(n int) int {
	//cpp， return n <= 3? n - 1 : pow(3, n / 3) * 4 / (4 - n % 3);
	if n <= 3 {
		return n - 1
	} else {
		return int(math.Pow(float64(3), float64(n/3)) * float64(4/(4-n%3)))
	}

	return -1
}
