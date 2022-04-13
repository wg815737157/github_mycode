package main

import "fmt"

func halfFind(a []int, target int) bool {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := low + (high-low)/2
		if a[mid] == target {
			return true
		}
		if a[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
func searchLeetCode(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[l] {
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

//
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		//在升序时异常
		if nums[mid] > target {
			if nums[mid] >= nums[0] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] <= nums[len(nums)-1] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

func main() {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(a, 0))
}
