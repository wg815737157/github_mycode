func permute(num []int) [][]int {
	numLen := len(num)
	if numLen == 1 {
		slice := num
		res := make([][]int, 0, 0)
		res = append(res, slice)
		return res
	}
	resNum := permute(num[1:])
	tmp := num[0]
	res := make([][]int, 0, 0)
	for _, v := range resNum {
		vlen := len(v)
		for i := 0; i < vlen; i++ {
			addArray := make([]int, 0, 0)
			addArray = append(addArray, v[:i]...)
			addArray = append(addArray, tmp)
			addArray = append(addArray, v[i:]...)
			res = append(res, addArray)
		}
		v = append(v, tmp)
		res = append(res, v)
	}
	return res
}
