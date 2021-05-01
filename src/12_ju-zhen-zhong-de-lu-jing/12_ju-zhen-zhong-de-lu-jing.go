package main

func exist(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if dfs(board, word, row, col, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, idx int) bool {
	//越界或者字符不相等
	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[0]) ||
		board[i][j] != word[idx] {
		return false
	}

	if idx == len(word)-1 {
		return true
	}

	//标记已被访问
	tmpCh := board[i][j]
	board[i][j] = '/'

	//上下左右检查一下
	check := (dfs(board, word, i+1, j, idx+1) || dfs(board, word, i-1, j, idx+1) ||
		dfs(board, word, i, j+1, idx+1) || dfs(board, word, i, j-1, idx+1))

	//恢复
	board[i][j] = tmpCh
	return check
}

// func main() {
// 	//board := {{"A","B","C","E"},{"S","F","C","S"},{"A","D","E","E"}}
// 	//word := "ABCCED"
// 	fmt.Println(exist(board, word))
// }
