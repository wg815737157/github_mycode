package main

import "fmt"

func multiply(a [][]int, b [][]int) [][]int {
	res := make([][]int, len(a))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(b[0]))
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for o := 0; o < len(b); o++ {
				res[i][j] += a[i][o] * b[o][j]
			}
		}
	}
	return res
}

func main() {

	a := [][]int{{1, 1}, {1, 1}}
	b := [][]int{{0, 1, 2}, {0, 1, 2}}
	fmt.Println(multiply(a, b))

}
