package main

import "fmt"

func permutation(num []int, n int, curIndex int, res *[][]int) {
	if curIndex == n-1 {
		fmt.Println(num)
		tmp := make([]int, len(num))
		copy(tmp, num)
		*res = append(*res, tmp)
		return
	}
	for i := curIndex; i < n; i++ {
		swap(num, curIndex, i)
		permutation(num, n, curIndex+1, res)
		swap(num, i, curIndex)
	}
}

func main() {
	a := []int{1, 2, 3, 4}
	n := len(a)
	res := [][]int{}
	permutation(a, n, 0, &res)
	fmt.Println(len(res), res)

}
