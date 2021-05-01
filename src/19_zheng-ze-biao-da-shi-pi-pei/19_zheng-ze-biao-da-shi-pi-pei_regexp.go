package main

import (
	"fmt"
	"regexp"
)

// 直接调regexp包，需要前后加^和$，否则可能p只匹配了s的中间一部分
func isMatch(s string, p string) bool {
	res, _ := regexp.MatchString("^"+p+"$", s)
	return res
}

func main() {
	s := "aa"
	p := "a"

	fmt.Println(isMatch(s, p)) //false
}
