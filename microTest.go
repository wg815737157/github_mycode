package main

import "fmt"

type Element struct {
	Color int //0 space 1 black 2white
}

//TODO
func initMatrix() [][]Element {
	return nil
}

func detect(matrix [][]Element) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	column := len(matrix[0])
	successBlackCount := 0
	successWhileCount := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j].Color == 0 {
				continue
			}
			//black
			if matrix[i][j].Color == 1 {
				// 判断行
				n := j + 1
				for ; n < j+5 && n < column; n++ {
					if matrix[i][n].Color == 1 && n == j+4 {
						successBlackCount++
						break
					}
					if matrix[i][n].Color == 1 {
						continue
					}
					break
				}
				// 判断列
				n = i + 1
				for ; n < i+5 && n < row; n++ {
					if matrix[n][j].Color == 1 && n == i+4 {
						successWhileCount++
						break

					}
					if matrix[n][j].Color == 1 {
						continue
					}
					break
				}

				//判断左后斜着
				m := i + 1
				n = j - 1
				count := 0
				for ; 0 <= n && m < row; count++ {
					if matrix[m][n].Color == 1 && count == 4 {
						return true
					}
					if matrix[m][n].Color == 1 {
						n--
						m++
						continue
					}
					break
				}
				// 判断右前斜着

			}
			// while

		}
	}

	//if successWhileCount >= 1 && successBlackCount == 0 {
	//	fmt.Println("白")
	//}
	if successWhileCount >= 2 || successBlackCount >= 2 ||(successWhileCount == 1 && successBlackCount == 1){
		fmt.Println("非法")
	}
}
func main() {
	matrix := initMatrix()
	fmt.Println(detect(matrix))
}
