package main

import "strconv"

func findNthDigit(n int) int {
	digit := 1
	var start int32 = 1
	var count int32 = 9
	for n > int(count) {
		n -= int(count)
		digit += 1
		start *= 10
		count = int32(digit) * start * 9
	}
	var num int32 = start + (n-1)/digit
	return strconv.Itoa()
}
