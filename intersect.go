package main

import "fmt"

func intersect(nums1, nums2 []int) []int {
	//set1 := make(map[int]int)
	set2 := make(map[int]int)

	ret := []int{}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, v := range nums2 {
		set2[v]++
	}
	for _, v := range nums1 {
		if _, ok := set2[v]; ok && set2[v] > 0 {
			set2[v]--
			ret = append(ret, v)
		}
	}
	return ret
}

func main() {
	a := []int{1, 2, 2, 3, 4, 5}
	b := []int{1, 2, 2, 2, 4}
	fmt.Println(intersect(a, b))
}
