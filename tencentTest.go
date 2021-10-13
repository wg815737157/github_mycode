package main

import (
	"fmt"
)

func mergeArray(arrayA, arrayB []int) []int {
	res := []int{}
	i, j := 0, 0
	for ; i < len(arrayA); i++ {
		for ; j < len(arrayB); j++ {
			//arrayB值 小于arrayA值
			if arrayB[j] < arrayA[i] {
				res = append(res, arrayB[j])
				continue
			}
			res = append(res, arrayA[i])
			break
		}
	}
	// 如果arrayB 最大值多于arrayA
	if j < len(arrayB) {
		return append(res, arrayB[j:]...)
	}
	return res
}

func findSortedMiddleOfTwoArray(arrayA, arrayB []int) int {
	res := mergeArray(arrayA, arrayB)
	fmt.Println("排序后的数组", res)
	if len(res)%2 == 0 {
		fmt.Println("数组下表", len(res)/2-1)
		return res[len(res)/2-1]
	}
	fmt.Println("数组下表", len(res)/2-1)
	return res[len(res)/2]
}

//func twoSum(inputArray []int) []int {
//	return nil
//
//}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//func initT() *TreeNode {
//	t := &TreeNode{1, nil, nil}
//	node1 := &TreeNode{4, nil, nil}
//	node2 := &TreeNode{6, nil, nil}
//	t.Left = &TreeNode{5, node1, node2}
//	node3 := &TreeNode{7, nil, nil}
//	node4 := &TreeNode{9, nil, nil}
//	t.Right = &TreeNode{8, node3, node4}
//	return t
//}

//func BigK(index int) int {
//
//}

func searchSingleValue(inputArray []int) int {
	res := 0
	for i := 0; i < len(inputArray); i++ {
		res = res ^ inputArray[i]
	}
	return res
}

func main() {
	inputArray := []int{6, 6, 2, 4, 3, 3, 1, 1}
	XOR := searchSingleValue(inputArray)
	var flag = 1
	for i := 0; i < 64; i++ {
		if flag&XOR == flag {
			break
		}
		flag = flag << 1
	}

	first, second := 0, 0
	for i := 0; i < len(inputArray); i++ {
		if flag&inputArray[i] == 0 {
			first ^= inputArray[i]
			continue
		}
		second ^= inputArray[i]
	}

	fmt.Println(first, second)

	//res := []int{}
	//t := initT()
	//if t == nil {
	//	return
	//}
	//q := &Queue{}
	//q.Add(t)
	//for q.Len() > 0 {
	//	currentLen := q.Len()
	//	for i := 0; i < currentLen; i++ {
	//		node := q.Remove().(*TreeNode)
	//		res = append(res, node.Val)
	//		if node.Left != nil {
	//			q.Add(node.Left)
	//		}
	//		if node.Right != nil {
	//			q.Add(node.Right)
	//		}
	//	}
	//}
	//fmt.Println(res)

	//l := list.New()
	//l.PushBack(1)
	//v, _ := l.Front().Value.(int)
	//e := l.Front()
	//l.Remove(e)
	//

	//fmt.Println(v, l.Len())

	//ji := make(chan int)
	//ou := make(chan int)
	//
	//intputInt := []int{1, 2, 3, 4, 5, 6}
	//
	//go func() {
	//	for {
	//		index := <-ji
	//		if index >= len(intputInt) {
	//			return
	//		}
	//		fmt.Println(intputInt[index])
	//		ou <- index + 1
	//	}
	//}()
	//ji <- 0
	//go func() {
	//	for {
	//		index := <-ou
	//		if index >= len(intputInt) {
	//			return
	//		}
	//		fmt.Println(intputInt[index])
	//		ji <- index + 1
	//	}
	//}()
	//
	//time.Sleep(time.Second)
}
