package main

import "fmt"

func solution2(aArray []int, b int, iterm []int, res *[][]int) {
	if b == 0 {
		*res = append(*res, iterm)
		return
	}
	for i := 0; i < len(aArray); i++ {
		curIterm := make([]int, len(iterm))
		copy(curIterm, iterm)
		curIterm = append(curIterm, aArray[i])
		tmpArray := make([]int, len(aArray))
		copy(tmpArray, aArray)
		s := append(tmpArray[:i], tmpArray[i+1:]...)
		solution2(s, b-1, curIterm, res)
	}
}

func solution3(a []int, num int, item []int, res *[][]int) {
	if len(item) == num {
		*res = append(*res, item)
		return
	}
	for i := 0; i < len(a); i++ {
		copyItem := make([]int, len(item))
		copy(copyItem, item)
		copyItem = append(copyItem, a[i])
		copyA := make([]int, len(a)-i-1)
		copy(copyA, a[i+1:])
		solution3(copyA, num, copyItem, res)
	}

}

func main() {
	a := 3
	aArray := []int{}
	for i := 1; i <= a; i++ {
		aArray = append(aArray, i)
	}
	res := [][]int{}
	solution3(aArray, 2, []int{}, &res)
	fmt.Println(res)
}
