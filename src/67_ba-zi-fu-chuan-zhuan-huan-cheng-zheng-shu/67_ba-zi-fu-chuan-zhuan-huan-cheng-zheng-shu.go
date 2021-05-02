func strToInt(str string) int {
	max := int(math.Pow(float64(2), float64(31))) - 1
	min := -int(math.Pow(float64(2), float64(31)))
	re := regexp.MustCompile(`^[+-]?\d+`)
	slice := re.FindAllString(strings.TrimSpace(str), -1)
	str = strings.Join(slice, "")
	res, _ := strconv.Atoi(str)
	if res < min {
		res = min
	} else if res > max {
		res = max
	}
	return res
}

//strconv.Atoi