func spiralOrder(matrix [][]int) []int {

	res := make([]int, 0, 0)

	left, top := 0, 0
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}

	right, bottom := len(matrix[0])-1, len(matrix)-1
	for left <= right && top <= bottom {
		//顶部遍历
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}

		//右侧遍历
		for j := top + 1; j <= bottom; j++ {
			res = append(res, matrix[j][right])
		}

		if left < right && top > bottom {
			//底部遍历
			for m := right - 1; m >= left; m-- {
				res = append(res, matrix[bottom][m])
			}
			//左侧遍历
			for n := bottom - 1; n >= top+1; n-- {
				res = append(res, matrix[n][left])
			}

		}

		left++
		top++
		bottom--
		right--
	}
	return res
}
