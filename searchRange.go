package main

func searchRangeLower(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func searchRangeHigher(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func searchRange(nums []int, target int) []int {
	l := searchRangeLower(nums, target)
	r := searchRangeHigher(nums, target)
	if l < len(nums) && r >= 0 && nums[l] == target && nums[r] == target {
		return []int{l, r}
	} else {
		return []int{-1, -1}
	}
}

func main() {

}
