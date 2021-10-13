package main

func removeDuplicates(nums []int) int {
	arrayLen := len(nums)
	if arrayLen == 0 {
		return 0
	}
	n := 0
	for i := 1; i < arrayLen; i++ {
		if nums[i-1] != nums[i] {
			nums[n] = nums[i-1]
			n++
			continue
		}
	}
	nums[n] = nums[arrayLen-1]
	return n + 1
}
