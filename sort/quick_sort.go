func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	len := len(nums)
	value := nums[len-1]
	low := 0
	high := len - 2
	for low < high {
		if nums[low] <= value {
			low++
			continue
		}
		if nums[high] >= value {
			high--
			continue
		}
		tmp := nums[low]
		nums[low] = nums[high]
		nums[high] = tmp
		low++
	}
	if nums[low] < value {
		tmp := nums[low+1]
		nums[low+1] = nums[len-1]
		nums[len-1] = tmp

		leftNums := quickSort(nums[0 : low+1])
		tmpNums := append(leftNums, nums[low+1])
		rightNums := quickSort(nums[low+2:])
		return append(tmpNums, rightNums...)
	} else {
		tmp := nums[low]
		nums[low] = nums[len-1]
		nums[len-1] = tmp

		leftNums := quickSort(nums[0:low])
		tmpNums := append(leftNums, nums[low])
		rightNums := quickSort(nums[low+1:])
		return append(tmpNums, rightNums...)
	}
}
