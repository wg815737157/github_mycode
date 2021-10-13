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

func isValid(s string) bool {
	return false
}

func main() {

	//mapBytes := map[byte]byte{
	//	']': '[',
	//}
	//fmt.Println(mapBytes[']'])

	bytes := []byte{}
	var i byte
	for i = -120; i < 120; i++ {
		bytes = append(bytes, i)
	}
	for i := 0; i < len(bytes); i++ {
		fmt.Println(i, bytes[i])
	}
	return
	//sliceA := make([][]int, 0, 0)
	//tmpSliceA := []int{6, 5, 4, 3, 2, 1}
	//tmpSliceB := []int{1, 2, 3, 4, 5, 6}
	//
	//sliceA = append(sliceA, tmpSliceA, tmpSliceB)
	//minArray(sliceA)
	//fmt.Println(flagKindWidth)
	//fmt.Println(1<<flagKindWidth)
	//fmt.Println(1<<flagKindWidth-1)
	//fmt.Println(flagKindMask)
	//
	//fmt.Println(Ptr)

	//nums := make([]int, 10, 20)

	//for _, num := range nums {
	//	fmt.Println(num)
	//}
	//var a Aer
	//
	//a = &StructA{}
	//B := &StructB{}

	//t := reflect.TypeOf(A)
	//if _, ok := t.(reflect.Type); ok {
	//	fmt.Println(ok)
	//}
	//fmt.Println(t.String())

	//v := reflect.ValueOf(a)
	//fmt.Println(v.Kind())
	//fmt.Println(v.String())

	//sliceA := make([]int, 10, 20)
	//sliceB := make([]int, 10, 20)

}
