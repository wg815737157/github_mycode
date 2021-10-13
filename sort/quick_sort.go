func quickSort(input []int) {
	if len(input) <= 1 {
		return
	}
	low := 0
	high := len(input) - 1
	pivot := input[0]
	for low < high {
		if input[high] >= pivot {
			high--
			continue
		}
		if input[low] <= pivot {
			low++
			continue
		}
		swap(input, low, high)
	}

	swap(input, low, 0)
	quickSort(input[:low])
	quickSort(input[low+1:])
	return
}
