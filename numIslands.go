package main

import "fmt"

func dfs(grid [][]byte, i int, j int) {
	grid[i][j] = 0
	if i-1 >= 0 {
		if grid[i-1][j] == '1' {
			dfs(grid, i-1, j)
		}
	}
	if i+1 < len(grid) {
		if grid[i+1][j] == '1' {
			dfs(grid, i+1, j)
		}
	}
	if j-1 >= 0 {
		if grid[i][j-1] == '1' {
			dfs(grid, i, j-1)
		}
	}
	if j+1 < len(grid[0]) {
		if grid[i][j+1] == '1' {
			dfs(grid, i, j+1)
		}
	}
}

func numIslands(grid [][]byte) int {
	row := len(grid)
	column := len(grid[0])
	num := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, i, j)
			}
		}
	}
	return num
}
func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid))
}
