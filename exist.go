package main

import (
	"fmt"
)

func dfs(board [][]byte, curRow int, curColumn int, s string, index int) bool {
	row := len(board)
	column := len(board[0])
	if index == len(s) {
		return true
	}

	if curRow < 0 || curColumn < 0 || curRow >= row || curColumn >= column {
		return false
	}

	if board[curRow][curColumn] != s[index] {
		return false
	}

	board[curRow][curColumn] = ' '

	res := dfs(board, curRow+1, curColumn, s, index+1) ||
		dfs(board, curRow-1, curColumn, s, index+1) ||
		dfs(board, curRow, curColumn+1, s, index+1) ||
		dfs(board, curRow, curColumn-1, s, index+1)

	board[curRow][curColumn] = s[index]

	return res
}

//[["A","B","C","E"],["S","F","E","S"],["A","D","E","E"]]
//"ABCESEEEFS"

func exist(board [][]byte, s string) bool {
	row := len(board)
	column := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(board, i, j, s, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	a := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	//b := make([][]byte, len(a))
	//for i := 0; i < len(a); i++ {
	//	b[i] = make([]byte, len(a[0]))
	//}
	//copy(b, a)
	//fmt.Printf("%p \n", a[0])
	//fmt.Printf("%p \n", b[0])

	//
	fmt.Println(exist(a, "ABCESEEEFS"))
}
