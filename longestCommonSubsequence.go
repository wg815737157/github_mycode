package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abd", "bd"))
}

func longestCommonSubsequence(s1, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	matrix := make([][]int, l1+1)
	for i := range matrix {
		matrix[i] = make([]int, l2+1)
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = max(matrix[i-1][j], matrix[i][j-1])
			}
		}
	}
	return matrix[l1][l2]
}
