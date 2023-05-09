package main

import (
	"fmt"
)

type StructA struct {
	A int
}
type StructB struct {
	A int
}

type Kind uint

type flag uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Map
	Ptr

	flagKindWidth      = 5 // there are 27 kinds
	flagKindMask  flag = 1<<flagKindWidth - 1
)

type Aer interface {
	Say()
}

func (A *StructA) Say() {
	fmt.Println("StructA Say")
}

//func merge(num int, permuteList [][]int) [][]int {
//
//}

//func permute(list []int) [][]int {
//	if len(list) == 0 {
//		return nil
//	}
//	if len(list) == 1 {
//		res := make([][]int, 0, 0)
//		return append(res, list)
//	}
//	return merge(list[0], permute(list[1:]))
//}
func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
func sumFactorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return factorial(n) + sumFactorial(n-1)
	}
}

func minArray(sliceA [][]int) {
	rangeTime := len(sliceA)
	columnLen := len(sliceA[0])
	sliceMinRes := make([]int, columnLen, columnLen)
	for i := 0; i < columnLen; i++ {
		tmp := sliceA[0][i]
		for m := 0; m < rangeTime; m++ {
			if sliceA[m][i] < tmp {
				tmp = sliceA[m][i]
			}
		}
		sliceMinRes[i] = tmp
	}
	fmt.Println(sliceMinRes)
}

func popSort(a []int) {
	for i := len(a); i > 0; i-- {
		for j := 1; j < i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}

func bitmap(a []int) []int {
	bytes := 1000 / 8
	bitMap := make([]byte, bytes)
	res := []int{}
	for i := 0; i < len(a); i++ {
		h := a[i] / 8
		l := 1 << (a[i] % 8)
		//存在
		if bitMap[h]&byte(l) == byte(l) {
			res = append(res, a[i])
		} else {
			bitMap[h] = bitMap[h] | byte(l)
		}
	}
	return res
}

func main() {
	s := "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"
	fmt.Println(len(s))

	//a := []int{0, 0, 0, 1, 2, 3, 4, 5, 5, 6, 9, 10, 100, 10, 9}
	//fmt.Println(bitmap(a))
}
