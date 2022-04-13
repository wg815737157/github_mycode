package main

import "fmt"

func dfs(grid [][]int, i, j int) int {
	res := 0
	if grid[i][j] > 0 {
		res++
		grid[i][j] = -1
	}
	if j-1 > 0 && grid[i][j-1] > 0 {
		res += dfs(grid, i, j-1)
	}
	if j+1 < len(grid[0]) && grid[i][j+1] > 0 {
		res += dfs(grid, i, j+1)
	}
	if i+1 < len(grid) && grid[i+1][j] > 0 {
		res += dfs(grid, i+1, j)
	}

	if i-1 > 0 && grid[i-1][j] > 0 {
		res += dfs(grid, i-1, j)
	}
	return res
}

func maxAreaOfIsland(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	maxLength := -1
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] <= 0 {
				continue
			}
			curLen := dfs(grid, i, j)
			if curLen > maxLength {
				maxLength = curLen
			}
		}
	}
	return maxLength
}

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid))
}
