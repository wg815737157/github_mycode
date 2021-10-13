package main

import "fmt"

func stepMethod(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return stepMethod(m-1, n) + stepMethod(m, n-1)
}

func main() {

	item1 := make([]int, 7, 7)
	item2 := make([]int, 7, 7)
	item3 := make([]int, 7, 7)

	a := [][]int{item1, item2, item3}
	row := len(a)
	column := len(a[0])

	for m := 0; m < row; m++ {
		for n := 0; n < column; n++ {
			if m == 0 || n == 0 {
				a[m][n] = 1
				continue
			}
			a[m][n] = a[m-1][n] + a[m][n-1]
		}
	}
	fmt.Println(a[2][6])
}
