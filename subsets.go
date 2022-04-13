package main

import "fmt"

func subsets(a []int) [][]int {
	res := [][]int{{}}
	for i := 0; i < len(a); i++ {
		tmp := make([][]int, len(res))
		for j := 0; j < len(tmp); j++ {
			tmp[j] = make([]int, len(res[j]))
			copy(tmp[j], res[j])
		}
		for j := 0; j < len(res); j++ {
			res[j] = append(res[j], []int{a[i]}...)
		}
		res = append(res, tmp...)
	}
	return res
}

func subsets2(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i < 1<<len(nums); i++ {
		tmp := []int{}
		for n := 0; n < len(nums); n++ {
			if i>>n&1 > 0 {
				tmp = append(tmp, nums[n])
			}
		}
		res = append(res, tmp)
	}
	return res
}
func main() {

	fmt.Println(subsets2([]int{1, 2, 3}))
}
