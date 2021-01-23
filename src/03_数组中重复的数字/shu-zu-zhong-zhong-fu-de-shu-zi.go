package main

import(
	"fmt"
)

func main() {
	nums := {1,2,2,3,4,5,6}
	fmt.Println(findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
    var nummap = make(map[int]bool)
    for _,num := range nums {
        if !nummap[num] {
            nummap[num] = true
        } else {
            return num
        }
    }

    return -1
}
