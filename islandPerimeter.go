package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	column := len(grid[0])
	if column == 0 {
		return 0
	}

	land := 0
	neighbor := 0

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 0 {
				continue
			}
			land++
			//上
			if i-1 >= 0 && grid[i][j]&grid[i-1][j] == 1 {
				neighbor++
			}
			//下
			if i+1 < row && grid[i][j]&grid[i+1][j] == 1 {
				neighbor++
			}
			//左
			if j-1 >= 0 && grid[i][j]&grid[i][j-1] == 1 {
				neighbor++
			}
			//右
			if j+1 < column && grid[i][j]&grid[i][j+1] == 1 {
				neighbor++
			}
		}
	}
	return 4*land - neighbor
}

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(grid))
}
