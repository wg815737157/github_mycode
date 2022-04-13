package main

import "fmt"

func execute(a []int, res []int) bool {
	for i := 1; i < len(a)-1; i++ {
		res[i+1] = a[i] - res[i-1] - res[i]
		if res[i+1] >= 2 || res[i+1] < 0 {
			return false
		}
	}
	return true
}
func faceTest(a []int) [][]int {
	results := [][]int{}
	res := make([]int, len(a))
	if a[0] == 0 {
		res[0] = 0
		res[1] = 0
		if execute(a, res) {
			results = append(results, res)
		}
	}
	if a[0] == 2 {
		res[0] = 1
		res[1] = 1
		if execute(a, res) {
			results = append(results, res)
		}
	}
	if a[0] == 1 {
		res[0] = 1
		res[1] = 0
		if execute(a, res) {
			results = append(results, res)
		}
		res2 := make([]int, len(a))
		res2[0] = 0
		res2[1] = 1
		if execute(a, res2) {
			results = append(results, res2)
		}
	}

	return results
}
func main() {
	a := []int{1, 1, 0}
	fmt.Println(faceTest(a))
}
