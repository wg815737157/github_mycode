package main

import "fmt"

func mergeSortRecursion(nums []int, l, r int, tmp []int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1

	mergeSortRecursion(nums, l, mid, tmp)
	mergeSortRecursion(nums, mid+1, r, tmp)
	cnt := 0
	i, j := l, mid+1
	for ; i <= mid && j <= r; {
		if nums[i] < nums[j] {
			tmp[cnt] = nums[i]
			cnt++
			i++
		} else {
			tmp[cnt] = nums[j]
			cnt++
			j++
		}
	}
	for i <= mid {
		tmp[cnt] = nums[i]
		cnt++
		i++
	}
	for j <= r {
		tmp[cnt] = nums[j]
		cnt++
		j++
	}
	for o := 0; o < r-l+1; o++ {
		nums[l+o] = tmp[o]
	}
}
func sortArray(nums []int) []int {
	tmp := make([]int, len(nums))
	l, r := 0, len(tmp)-1
	mergeSortRecursion(nums, l, r, tmp)
	return nums
}

func main() {
	tmp := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray(tmp))
}
