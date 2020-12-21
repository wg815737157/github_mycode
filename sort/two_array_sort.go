func O_n(sliceA []int, sliceB []int) []int {
	sliceALen := len(sliceA)
	sliceBLen := len(sliceB)
	len := sliceALen + sliceBLen
	resSlice := make([]int, len)
	if sliceBLen == 0 && sliceALen == 0 {
		return resSlice
	}
	n, m := 0, 0
	for i := 0; i < len; i++ {
		if n == sliceALen && m < sliceBLen {
			resSlice[i] = sliceB[m]
			m++
		} else if n < sliceALen && m == sliceBLen {
			resSlice[i] = sliceA[n]
			n++
		} else if sliceA[n] < sliceB[m] {
			resSlice[i] = sliceA[n]
			n++
		} else {
			resSlice[i] = sliceB[m]
			m++
		}
	}
	return resSlice
}

func main() {
	sliceA := []int{3, 4, 5, 6, 7, 10, 11}
	sliceB := []int{1, 4, 8, 99, 100, 999}
	fmt.Println(O_n(sliceA, sliceB))
}
