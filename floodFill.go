package main

import (
	"fmt"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	row := len(image)
	column := len(image[0])
	curValue := image[sr][sc]
	if curValue == newColor {
		return image
	}

	q := &Queue{}
	q.Add([]int{sr, sc})
	for q.Len() != 0 {
		s := q.Remove().([]int)
		sr = s[0]
		sc = s[1]
		if image[sr][sc] != newColor {
			image[sr][sc] = newColor
		}

		if sr-1 >= 0 && image[sr-1][sc] == curValue {
			q.Add([]int{sr - 1, sc})
		}
		if sr+1 < row && image[sr+1][sc] == curValue {
			q.Add([]int{sr + 1, sc})
		}
		if sc-1 >= 0 && image[sr][sc-1] == curValue {
			q.Add([]int{sr, sc - 1})
		}
		if sc+1 < column && image[sr][sc+1] == curValue {
			q.Add([]int{sr, sc + 1})
		}
	}
	return image

}
func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	//[[2, 2, 2], [2, 2, 0], [2, 0, 1]]
	fmt.Println(floodFill(image, 1, 1, 2))
}
