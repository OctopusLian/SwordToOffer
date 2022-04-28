/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-28 22:03:18
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-28 22:04:34
 */
package main

func findRepeatNumber(nums []int) int {
	var nummap = make(map[int]bool)
	for _, num := range nums {
		if !nummap[num] {
			nummap[num] = true
		} else {
			return num
		}
	}

	return -1
}
