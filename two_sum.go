func twoSum(nums []int, target int) ([]interface{}, error) {
	if len(nums) < 2 {
		return nil, errors.New("length too short")
	}
	low := 0
	high := len(nums) - 1
	res := make([]interface{}, 0, 0)
	for low < high {
		if nums[low]+nums[high] == target {
			tmp := make([]int, 0, 0)
			//fmt.Println(nums[low], nums[high])
			tmp = append(tmp, low, high)
			res = append(res, tmp)
			low++
			high--
			continue
		}
		if nums[low]+nums[high] > target {
			high--
			continue
		}
		if nums[low]+nums[high] < target {
			low++
			continue
		}
	}
	return res, nil
}
