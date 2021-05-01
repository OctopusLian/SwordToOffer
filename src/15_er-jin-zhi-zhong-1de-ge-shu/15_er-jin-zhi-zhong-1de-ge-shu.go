func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res += int(num & 1) //uint32转int
		num >>= 1
	}

	return res
}