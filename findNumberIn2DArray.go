package main

import "fmt"

//
//现有矩阵 matrix 如下：
//
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//给定 target=5，返回true。
//
//给定 target = 20，返回 false。
func findNumberIn2DArray(a [][]int, target int) bool {
	line := len(a)
	if line == 0 {
		return false
	}
	column := len(a[0])
	if a[0][column-1] == target {
		return true
	}
	if target < a[0][column-1] {
		for i := 0; i < line; i++ {
			a[i] = a[i][:column-1]
		}
		return findNumberIn2DArray(a, target)
	}
	if target > a[0][column-1] {
		return findNumberIn2DArray(a[1:], target)
	}
	return false
}
func main() {
	array2D := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 20
	fmt.Println(findNumberIn2DArray(array2D, target))
}
