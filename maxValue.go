package main

import "fmt"

//
//输入:
//[
//[1,3,1],
//[1,5,1],
//[4,2,1]
//]
//输出: 12
//解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
func maxValue(matrix [][]int) int {
	row := len(matrix)
	column := len(matrix[0])
	if column == 0 {
		return 0
	}
	res := make([][]int, row)
	for i := 0; i < row; i++ {
		res[i] = make([]int, column)
	}
	res[0][0] = matrix[0][0]
	for i := 1; i < row; i++ {
		res[i][0] = res[i-1][0] + matrix[i][0]
	}
	for j := 1; j < column; j++ {
		res[0][j] = res[0][j-1] + matrix[0][j]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			res[i][j] = twoMaxNum(res[i][j-1], res[i-1][j]) + matrix[i][j]
		}
	}
	return res[row-1][column-1]
}

func main() {
	arrayA := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(maxValue(arrayA))
}
