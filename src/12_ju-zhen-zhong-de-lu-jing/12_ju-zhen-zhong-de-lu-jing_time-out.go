package main

import "fmt"

//为什么会超时？
func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//board2 := [][]byte{{'a'}, {'a'}}
	//board3 := [][]byte{{'a', 'b'}}
	//board4 := [][]byte{{'a', 'b'}, {'c', 'd'}}
	//board5 := [][]byte{{'c', 'a', 'a'}, {'a', 'a', 'a'}, {'b', 'c', 'd'}}
	//board6 := [][]byte{{'a'}}
	word := "ABCCED"
	//word2 := "aaa"
	//word3 := "ba"
	//word4 := "cdba"
	//word5 := "aab"
	//word6 := "ab"
	fmt.Println(exist(board, word))
}
func exist(board [][]byte, word string) bool {
	//异常一[["a"]]"ab"
	//内部多次调用，但是全局变量未重新赋值，手动赋值尝试
	//异常二[["a"]]"a"
	//边界异常
	//异常三[["a"]，["a"]]"a"
	//在异常二的基础上要多考虑，尽量让简单运算或者较优条件先执行
	//异常四[["a"]，["a"]]"aaa"
	//修改终止条件
	//异常五[['A', 'B']]"BA"
	//可重复
	c := 0
	g := 0
	for c != len(board[0]) && g != len(board) {
		n := 0
		var nboard [200][200]int
		if board[g][c] == word[n] {
			if n >= len(word) || dfs(board, nboard, g, c, word, n) {
				y = 0
				return true
			}
		}
		if c == len(board[0])-1 {
			g++
			c = 0
		} else {
			c++
		}
	}
	return false
}

var y int

func dfs(board [][]byte, nboard [200][200]int, g int, c int, word string, n int) bool {
	if n == len(word) {
		y = 1
		return true
	}
	if g < 0 || c < 0 || g == len(board) || c == len(board[0]) {
		return false
	}
	if board[g][c] == word[n] && nboard[g][c] != 1 {
		nboard[g][c] = 1
		if dfs(board, nboard, g-1, c, word, n+1) ||
			dfs(board, nboard, g+1, c, word, n+1) ||
			dfs(board, nboard, g, c-1, word, n+1) ||
			dfs(board, nboard, g, c+1, word, n+1) {
			y = 1
		}
	}
	if y == 1 {
		return true
	}
	return false
}

//链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/solution/yi-zhi-chao-shi-xu-yao-hou-xu-yan-jiu-by-p94j/
