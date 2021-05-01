package main

import "fmt"

func fib(n int) int {
	a, b := 0, 1
	var sum int
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}

	return a
}

func main() {
	fmt.Println(fib(0))
}
