package main

import "fmt"

func sumNums(n int) int {
	if n <= 0 {
		return n
	}
	matrix := make([]int, n+1)
	matrix[0] = 0

	for i := 1; i < n+1; i++ {
		matrix[i] = matrix[i-1] + i
	}
	return matrix[n]
}

func main() {

	i := 9
	fmt.Println(sumNums(i))

}
